package logic

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"ywadmin-v3/common/tool"
	"ywadmin-v3/service/monitor/rpc/internal/svc"
	"ywadmin-v3/service/monitor/rpc/monitorclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectReportLogic {
	return &SelectReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectReportLogic) SelectReport(in *monitorclient.SelectReportReq) (*monitorclient.SelectReportResp, error) {
	i, err := l.svcCtx.ReportStreamMinuteModel.Count(l.ctx, in.AssetId, in.ReportTime, in.AssetIp)
	if err != nil {
		return nil, err
	}
	//fmt.Println(i, "     ip---->", in.AssetIp, "  ", in.AssetId)

	if i == 0 {
		//创建当天的ip文件
		f, split := createDirandFile("./date")
		if !tool.StrInArr(in.AssetIp, split) {
			//发送报警，并写文件
			_, err = l.svcCtx.ReportStreamMinuteModel.InsertReport(l.ctx, in.AssetIp, in.Remark)
			if err != nil {
				return nil, err
			}
			f.WriteString(in.AssetIp + "\n")
		}
	}

	return &monitorclient.SelectReportResp{
		Count: i,
	}, nil
}

func createDirandFile(dir string) (f *os.File, split []string) {
	exists, _ := PathExists(dir)
	if !exists {
		os.MkdirAll(dir, 0666)
	}
	//当天的日期
	file := fmt.Sprintf("%s/%s.txt", dir, time.Now().Format("20060102"))
	existsFile, _ := PathExists(file)
	if !existsFile {
		os.Create(file)
	}

	//读取今天已发送的消息
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	split = strings.Split(strings.TrimSpace(string(content)), "\n")
	f, err = os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
