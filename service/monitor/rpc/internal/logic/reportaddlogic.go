package logic

import (
	"bytes"
	"context"
	"time"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/monitor/rpc/internal/svc"
	"ywadmin-v3/service/monitor/rpc/monitorclient"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/util/gconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReportAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportAddLogic {
	return &ReportAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReportAddLogic) ReportAdd(in *monitorclient.ReportAddReq) (*monitorclient.ReportAddResp, error) {
	var err error
	now := time.Now().Unix()
	//加密校验
	if Sign(in.Ts, in.Sn, l.svcCtx.Config.KeyDependency.AppKey, l.svcCtx.Config.KeyDependency.SecretCode) == false {
		return nil, xerr.NewErrMsg("校验加密失败")
	}
	//判断超时(10分钟超时)
	if (now-gconv.Int64(in.Ts)) > 300 || (now-gconv.Int64(in.Ts)) < -300 {
		return nil, xerr.NewErrMsg("加密字符串非当前时间10分钟内，请检查系统时间，然后重新生成加密字符")
	}
	if in.ReportType == "monitor" {
		_, err = l.svcCtx.ReportStreamMinuteModel.Insert(l.ctx, in)
	} else if in.ReportType == "machine" {
		err = l.svcCtx.ReportStreamMinuteModel.UpdateAssetHwInfo(l.ctx, in)
	} else {
		return nil, xerr.NewErrMsg("上报类型错误")
	}
	if err != nil {
		return nil, xerr.NewErrMsg("增加数据失败，原因：" + err.Error())
	}

	return &monitorclient.ReportAddResp{}, nil
}

func Sign(ts, sign, appKey, secretCode string) bool {
	b := bytes.Buffer{}
	b.WriteString("appKey=")
	b.WriteString(appKey)
	b.WriteString("&appSecret=")
	b.WriteString(secretCode)
	b.WriteString("&ts=")
	b.WriteString(ts)
	//fmt.Println(b.String())
	//fmt.Println("md5:", gmd5.MustEncryptString(b.String()))
	if gmd5.MustEncryptString(b.String()) == sign {
		return true
	}
	return false
}
