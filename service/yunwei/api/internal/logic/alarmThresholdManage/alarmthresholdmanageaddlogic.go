package alarmThresholdManage

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmThresholdManageAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlarmThresholdManageAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmThresholdManageAddLogic {
	return &AlarmThresholdManageAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlarmThresholdManageAddLogic) AlarmThresholdManageAdd(req *types.AddAlarmThresholdManageReq) error {
	var tmp yunweiclient.AlarmThresholdManageCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.AlarmThresholdManageAdd(l.ctx, &yunweiclient.AddAlarmThresholdManageReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil

}
