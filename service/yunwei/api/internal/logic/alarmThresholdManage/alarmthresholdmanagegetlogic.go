package alarmThresholdManage

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmThresholdManageGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlarmThresholdManageGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmThresholdManageGetLogic {
	return &AlarmThresholdManageGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlarmThresholdManageGetLogic) AlarmThresholdManageGet(req *types.GetAlarmThresholdManageReq) (resp *types.ListAlarmThresholdManageData, err error) {

	get, err := l.svcCtx.YunWeiRpc.AlarmThresholdManageGet(l.ctx, &yunweiclient.GetAlarmThresholdManageReq{
		AlarmThresholdManageId:    req.AlarmThresholdManageId,
		AlarmThresholdManageTypes: req.AlarmThresholdManageTypes,
	})
	if err != nil {
		return nil, err
	}
	var tmp types.ListAlarmThresholdManageData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
