package alarmThresholdManage

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunwei"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmThresholdManageUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlarmThresholdManageUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmThresholdManageUpdateLogic {
	return &AlarmThresholdManageUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlarmThresholdManageUpdateLogic) AlarmThresholdManageUpdate(req *types.UpdateAlarmThresholdManageReq) error {
	var tmp yunwei.AlarmThresholdManageCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("更新拷贝数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.AlarmThresholdManageUpdate(l.ctx, &yunweiclient.UpdateAlarmThresholdManageReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
