package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"ywadmin-v3/common/tasksFunc"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/jobs"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskGetFormatJsonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskGetFormatJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskGetFormatJsonLogic {
	return &TaskGetFormatJsonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskGetFormatJsonLogic) TaskGetFormatJson(in *yunweiclient.TaskGetFormatJsonReq) (*yunweiclient.TaskGetFormatJsonResp, error) {

	//获取文件路径
	file := time.Now().Format("20060102150405")
	file = fmt.Sprintf("%s_%s.txt", in.Game, file)
	dir, _ := filepath.Abs(filepath.Dir(l.svcCtx.Config.Scripts.MaintainFilePath))
	filePath := fmt.Sprintf("%s/format_files/", dir)
	if !tool.IsExist(filePath) {
		os.MkdirAll(filePath, os.ModePerm)
	}
	err := ioutil.WriteFile(filePath+"/"+file, []byte(in.Content), 0666)
	if err != nil {
		return nil, xerr.NewErrMsg("写入文件失败，原因：" + err.Error())
	}
	var uname string
	if in.Uname == "" {
		if md, ok := metadata.FromIncomingContext(l.ctx); ok {
			uname = md.Get("uname")[0]
			uname, _ = url.QueryUnescape(uname)
		}
	} else {
		uname = in.Uname
	}

	cmds := fmt.Sprintf("source /etc/profile;cd %s;sh cmd_entrance.sh -g %s -f %s/%s -op format -t TWEB -u %s", dir, in.Game, filePath, file, uname)
	jobOut := jobs.NewCommandJob(l.svcCtx, l.ctx, 24*time.Hour, cmds, "other", 1)
	if jobOut.ErrMsg != "" || !jobOut.IsOk {
		return nil, xerr.NewErrMsg(in.Game + "执行命令失败，原因：" + jobOut.ErrMsg + cmds)
	}
	if jobOut.IsTimeout {
		return nil, xerr.NewErrMsg(in.Game + "执行命令超时")
	}

	hotFile := fmt.Sprintf("%s/tmp/format/%s/batch_file.sh", dir, in.Game)
	hot := FormMatHot(hotFile, l.svcCtx, l.ctx)
	var tmp []*yunweiclient.OperationListM
	err = copier.Copy(&tmp, hot)
	if err != nil {
		return nil, xerr.NewErrMsg("复制格式化数据失败，原因：" + err.Error())
	}

	return &yunweiclient.TaskGetFormatJsonResp{
		OperationListM: tmp,
	}, nil
}

func sysOperationServiceGetList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*OperationGroupJson, error) {
	result, err := svcCtx.AdminRpc.MenuOperationList(ctx, &adminclient.MenuOperationListReq{})
	if err != nil {
		return nil, err
	}
	list := make([]*OperationGroupJson, 0)
	for _, v := range result.MenuOperationListData {
		tmp := new(OperationGroupJson)
		tmp.Label = v.ParentName
		tmpinfo := make([]OperationList, 0)

		nameSplit := strings.Split(v.Name, ",")
		urlSplit := strings.Split(v.Url, ",")

		for i := 0; i < len(nameSplit); i++ {
			tasksFunc.MapsOperation[urlSplit[i]] = nameSplit[i]
			tmpinfo = append(tmpinfo, OperationList{
				Label: nameSplit[i],
				Value: urlSplit[i],
			})
		}
		tmp.CanOperationList = tmpinfo
		if len(tmpinfo) != 0 {
			list = append(list, tmp)
		}
	}
	logx.Error(tasksFunc.MapsOperation)
	return list, nil
}

func FormMatHot(dir string, svcCtx *svc.ServiceContext, ctx context.Context) []*OperationListM {
	//a = `echo "任务1";sh /root/zhangqingxian/yunwei_new/maintain_game/main.sh -g bkzj -o server_hot -t c0b933ce973f5d381d702f83993152c1_1635130894307666760 -u zhangqingxian -j '{'\\"'stable'\\"':'\\"'1'\\"','\\"'outerIp'\\"':'\\"'114.67.222.39'\\"','\\"'fileList'\\"':'\\"'lua/common/test.lua,lua/game/hero/resonancemanager.lua'\\"','\\"'maintainRange'\\"':'\\"'bkzjjh15,bkzjjh16,bkzjjh17,bkzjjh18,bkzjjh20,bkzjjh21,Bkzjjh15cross,Bkzjjh16cross,Bkzjjh17cross,Bkzjjh18cross,Bkzjjh20cross,Bkzjjh21cross'\\"'}'
	//echo "任务2";sh /root/zhangqingxian/yunwei_new/maintain_game/main.sh -g bkzj -o server_cmd -t 7546598e2a03159bf5191306c97aadfc_1635130894341819340 -u zhangqingxian -j '{'\\"'cmdList'\\"':'\\"'${CDP_HOME}/skynet/lua ${CDP_HOME}/lua/common/gmserver.lua ${CDP_HOME}/skynet/ updateconfig test'\\"','\\"'maintainRange'\\"':'\\"'bkzjjh15,bkzjjh16,bkzjjh17,bkzjjh18,bkzjjh20,bkzjjh21,Bkzjjh15cross,Bkzjjh16cross,Bkzjjh17cross,Bkzjjh18cross,Bkzjjh20cross,Bkzjjh21cross'\\"'}'`

	result, err := sysOperationServiceGetList(svcCtx, ctx)
	if err != nil {
		return nil
	}

	groupJson, err := json.Marshal(result)
	if err != nil {
		return nil
	}
	content, _ := ioutil.ReadFile(dir)
	lists := strings.Split(string(content), "\n")

	var operPre string
	count := 0
	tmps := make([][]interface{}, 0)
	for i := 0; i < len(lists)-1; i++ {
		//格式化每个任务操作类型
		compile, _ := regexp.Compile(`-o|-j`)
		tmp := compile.Split(lists[i], -1)
		operType := strings.Split(tmp[1], " ")
		listForm := strings.ReplaceAll(tmp[2], "'\\\"'", "\"")
		listForm = strings.TrimLeft(listForm, " ")
		listForm = trimQuotes(listForm)
		m := new(TaskCommonJson)
		err := json.Unmarshal([]byte(listForm), &m)
		if err != nil {
			return nil
		}
		s, b := grouping(groupJson, operType[1], operPre)
		m.Operation = operType[1]
		m.OperationCn = tasksFunc.MapsOperation[operType[1]]
		tmps = append(tmps, []interface{}{
			m, b,
		})
		if !b {
			count++
		}
		operPre = s

	}

	tt := make([]*OperationListM, count)
	cc := make([]*TaskCommonJson, 0)
	n := 0
	for i := 0; i < len(tmps); i++ {
		if tmps[i][1].(bool) {
			n++
			cc = append(cc, tmps[i][0].(*TaskCommonJson))
			tt[i-n] = &OperationListM{
				OperationListForm: cc,
			}
		} else {
			cc = make([]*TaskCommonJson, 0)
			cc = append(cc, tmps[i][0].(*TaskCommonJson))
			tt[i-n] = &OperationListM{
				OperationListForm: cc,
			}
		}
	}

	//fmt.Println(tt)
	//marshal, _ := json.Marshal(tt)
	//fmt.Println(string(marshal))
	return tt

}

func grouping(groupJs []byte, operCurent, operPre string) (string, bool) {
	var groups []OperationGroupJson
	err := json.Unmarshal(groupJs, &groups)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range groups {
		for _, v1 := range v.CanOperationList {
			if v1.Value == operCurent && v.Label != operPre {
				return v.Label, false
			}
			if v1.Value == operCurent && v.Label == operPre {
				return v.Label, true
			}
		}

	}
	return "", false

}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && c == '\'' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// 解析前端json
type OperationListM struct {
	OperationListForm []*TaskCommonJson `json:"operationListForm"`
}

// 解析统一结构-对应下面内容
type TaskCommonJson struct {
	//Id            int64  `json:"id,omitempty"`
	Operation      string `json:"operation,omitempty"`
	OperationCn    string `json:"operationCn,omitempty"`
	Stable         string `json:"stable,omitempty"`
	OuterIp        string `json:"outerIp,omitempty"`
	DbUpdate       string `json:"dbUpdate,omitempty"`
	FileList       string `json:"fileList,omitempty"`
	MaintainRange  string `json:"maintainRange,omitempty"`
	CmdList        string `json:"cmdList,omitempty"`
	DbType         string `json:"dbType,omitempty"`
	PlatName       string `json:"platName,omitempty"`
	CheckSt        string `json:"checkSt,omitempty"`
	SQLCmd         string `json:"SQLCmd,omitempty"`
	Merge          string `json:"merge,omitempty"`
	ExportFileName string `json:"exportFileName,omitempty"`
	InitSetTime    string `json:"initSetTime,omitempty"`
	ExecuteSQL     string `json:"executeSQL,omitempty"`
	ExecuteFlag    string `json:"executeFlag,omitempty"`
}

// 操作权限分组展示
type OperationGroupJson struct {
	Label            string          `json:"label,omitempty"`
	CanOperationList []OperationList `json:"canOperationList,omitempty"`
}

type OperationList struct {
	Value string `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
}
