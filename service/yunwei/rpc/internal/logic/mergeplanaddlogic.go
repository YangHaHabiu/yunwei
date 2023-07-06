package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergePlanAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMergePlanAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanAddLogic {
	return &MergePlanAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MaintainPlan Rpc End
func (l *MergePlanAddLogic) MergePlanAdd(in *yunweiclient.AddMergePlanReq) (*yunweiclient.MergePlanCommonResp, error) {
	err := l.svcCtx.MergePlanModel.TransactInsert(l.ctx, in.MergePlanData)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}
	return &yunweiclient.MergePlanCommonResp{}, nil
}
