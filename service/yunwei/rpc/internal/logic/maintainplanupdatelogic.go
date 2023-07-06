package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/metadata"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMaintainPlanUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanUpdateLogic {
	return &MaintainPlanUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MaintainPlanUpdateLogic) MaintainPlanUpdate(in *yunweiclient.UpdateMaintainPlanReq) (*yunweiclient.MaintainPlanCommonResp, error) {
	var tmp model.MaintainPlan
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝更新数据失败，原因：" + err.Error())
	}
	var uid string
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		uid = md.Get("uid")[0]
	}
	tmp.UpdateBy = uid
	one, err := l.svcCtx.MaintainPlanModel.FindOne(l.ctx, tmp.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败，原因：" + err.Error())
	}
	if one.TaskId != -1 {
		return nil, xerr.NewErrMsg("存在运行中的任务，不允许编辑")
	}

	err = l.svcCtx.MaintainPlanModel.Update(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}
	return &yunweiclient.MaintainPlanCommonResp{}, nil
}
