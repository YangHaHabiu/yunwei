package alarmThresholdManage

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmThresholdManageDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlarmThresholdManageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmThresholdManageDeleteLogic {
	return &AlarmThresholdManageDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlarmThresholdManageDeleteLogic) AlarmThresholdManageDelete(req *types.DeleteAlarmThresholdManageReq) error {
	_, err := l.svcCtx.YunWeiRpc.AlarmThresholdManageDelete(l.ctx, &yunweiclient.DeleteAlarmThresholdManageReq{AlarmThresholdManageId: req.AlarmThresholdManageId})
	if err != nil {
		return err
	}
	return nil
}
