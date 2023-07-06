package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmThresholdManageAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlarmThresholdManageAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmThresholdManageAddLogic {
	return &AlarmThresholdManageAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AlarmThresholdManage Rpc Start
func (l *AlarmThresholdManageAddLogic) AlarmThresholdManageAdd(in *yunweiclient.AddAlarmThresholdManageReq) (*yunweiclient.AlarmThresholdManageCommonResp, error) {
	var tmp model.AlarmThresholdManage
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.AlarmThresholdManageModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}

	return &yunweiclient.AlarmThresholdManageCommonResp{}, nil
}
