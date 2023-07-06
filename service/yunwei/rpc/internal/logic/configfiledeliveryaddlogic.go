package logic

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xcmd"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/common/xsshClient"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/gogf/gf/util/gconv"
	"github.com/panjf2000/ants/v2"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileDeliveryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryAddLogic {
	return &ConfigFileDeliveryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type projectData struct {
	ProjectEn string
	List      []*yunweiclient.AddConfigFileDeliveryData
}

// ConfigFile Rpc End
func (l *ConfigFileDeliveryAddLogic) ConfigFileDeliveryAdd(in *yunweiclient.AddConfigFileDeliveryReq) (*yunweiclient.AddConfigFileDeliveryResp, error) {
	//检查锁文件
	for _, v := range in.ConfigFileData {
		lockpath := filepath.Join(l.svcCtx.Config.LockFilePath, v.ProjectEn)
		if !tool.IsExist(lockpath) {
			tool.CreateMutiDir(lockpath)
		}
		if tool.IsExist(filepath.Join(lockpath, "lock")) {
			return nil, xerr.NewErrMsg(fmt.Sprintf("%s存在锁文件，请检查，路径：%s", v.ProjectEn, lockpath))
		}
	}

	//开始批量下发

	for _, v := range in.ConfigFileData {
		//创建锁文件
		fileName := filepath.Join(l.svcCtx.Config.LockFilePath, v.ProjectEn, "lock")
		_, err := os.Create(fileName)
		if err != nil {
			return nil, xerr.NewErrMsg("创建锁文件失败，原因：" + err.Error())
		}
		//同步下脚本
		configCenterPath := fmt.Sprintf("source /etc/profile;bash %s/auto_create_env.sh %s", l.svcCtx.Config.ConfigCenterPath, v.ProjectEn)
		job := xcmd.NewCommandJob(1*time.Minute, configCenterPath)
		if job.ErrMsg != "" || !job.IsOk {
			os.Remove(fileName)
			return nil, xerr.NewErrMsg("生成配置文件失败，原因：" + job.ErrMsg + configCenterPath)
		}
		if job.IsTimeout {
			os.Remove(fileName)
			return nil, xerr.NewErrMsg("执行命令超时")
		}
		//var configEnv yunweiclient.AddConfigFileDeliveryData
		//copier.Copy(&configEnv, v.List[0])
		//configEnv.ConfigName = "ConfigEnv.sh"
		//v.List = append(v.List, &configEnv)
		//fmt.Println(v.List)

		//异步执行下发任务
		go asyncExec(l.svcCtx.Config.KeyFullPath,
			filepath.Join(l.svcCtx.Config.ConfigCenterPath, v.ProjectEn),
			filepath.Join(l.svcCtx.Config.LockFilePath, v.ProjectEn, "lock"),
			v.List,
			l.svcCtx.Config.ConfigMngThreads,
		)
	}

	//修改（添加）数据
	l.Logger.Error("------------>", in)
	err := l.svcCtx.ConfigMngLogModel.TransactAdd(l.ctx, in, l.svcCtx.Config.LockFilePath)

	if err != nil {
		return nil, xerr.NewErrMsg("批量下发配置文件失败，原因：" + err.Error())
	}

	return &yunweiclient.AddConfigFileDeliveryResp{}, nil
}

//批量执行-控制线程数100
func asyncExec(key, soucePath, lockFile string, in []*yunweiclient.AddConfigFileDeliveryData, threadNums int) {
	//p, _ := ants.NewPool(threadNums)

	var wg sync.WaitGroup
	length := len(in)

	p, _ := ants.NewPoolWithFunc(threadNums, func(i interface{}) {
		myFunc(i, in, key, soucePath)
		wg.Done()
	})
	defer p.Release()

	for i := 0; i < length; i++ {
		wg.Add(1)
		err := p.Invoke(int32(i))

		if err != nil {
			os.Remove(lockFile)
		}
	}
	wg.Wait()
	//删除锁文件
	os.Remove(lockFile)
}

func myFunc(i interface{}, in []*yunweiclient.AddConfigFileDeliveryData, key, soucePath string) {
	x := gconv.Int(i)
	batchScpFile(in[x].SshIp,
		key,
		in[x].SshPort,
		filepath.Join(soucePath, in[x].ConfigName),
		filepath.Join(in[x].DestPath, in[x].ConfigName))
}

//远程分发文件
func batchScpFile(ip, key string, port int64, soureFile, destFile string) error {
	sshClient, err := xsshClient.NewSSHClient(ip, gconv.Int(port), xsshClient.AuthConfig{User: "root", KeyFile: key})
	if err != nil {
		fmt.Println("远程连接：", err)
		return err
	}
	defer sshClient.Close()

	sshClient.Exec(fmt.Sprintf("[ -d %s ] || mkdir -p %s", filepath.Dir(destFile), filepath.Dir(destFile)))
	_, err = sshClient.Upload(soureFile, destFile)
	if err != nil {
		fmt.Println("传送：", err)
		return err
	}
	sshClient.Exec(fmt.Sprintf("chmod +x %s", destFile))

	return nil
}
