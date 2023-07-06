package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMaintainPlanDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanDeleteLogic {
	return &MaintainPlanDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MaintainPlanDeleteLogic) MaintainPlanDelete(in *yunweiclient.DeleteMaintainPlanReq) (*yunweiclient.MaintainPlanCommonResp, error) {
	one, err := l.svcCtx.MaintainPlanModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败，原因：" + err.Error())
	}
	if one.TaskId != 0 {
		return nil, xerr.NewErrMsg("存在运行中的任务，不允许删除")
	}
	err = l.svcCtx.MaintainPlanModel.DeleteSoft(l.ctx, in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &yunweiclient.MaintainPlanCommonResp{}, nil
}
