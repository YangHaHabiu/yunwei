package logic

import (
	"context"
	"fmt"
	"time"
	"ywadmin-v3/common/xcmd"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryGetLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileDeliveryGetLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryGetLogLogic {
	return &ConfigFileDeliveryGetLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigFileDeliveryGetLogLogic) ConfigFileDeliveryGetLog(in *yunweiclient.GetConfigFileDeliveryLogReq) (*yunweiclient.GetConfigFileDeliveryLogResp, error) {

	execCmd := fmt.Sprintf("source /etc/profile;bash %s/%s", l.svcCtx.Config.ConfigCenterPath, "insert_filelist.sh")
	job := xcmd.NewCommandJob(1*time.Minute, execCmd)
	if !job.IsOk || job.ErrMsg != "" {
		return nil, xerr.NewErrMsg("执行命令失败，原因：" + job.ErrMsg)
	}
	if job.IsTimeout {
		return nil, xerr.NewErrMsg("执行命令超时")
	}
	return &yunweiclient.GetConfigFileDeliveryLogResp{}, nil
}
