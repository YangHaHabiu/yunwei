package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmThresholdManageDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlarmThresholdManageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmThresholdManageDeleteLogic {
	return &AlarmThresholdManageDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlarmThresholdManageDeleteLogic) AlarmThresholdManageDelete(in *yunweiclient.DeleteAlarmThresholdManageReq) (*yunweiclient.AlarmThresholdManageCommonResp, error) {
	err := l.svcCtx.AlarmThresholdManageModel.DeleteSoft(l.ctx, in.AlarmThresholdManageId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}

	return &yunweiclient.AlarmThresholdManageCommonResp{}, nil
}
