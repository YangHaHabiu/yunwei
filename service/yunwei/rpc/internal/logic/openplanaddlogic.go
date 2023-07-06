package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenPlanAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenPlanAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanAddLogic {
	return &OpenPlanAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MergePlan Rpc End
func (l *OpenPlanAddLogic) OpenPlanAdd(in *yunweiclient.AddOpenPlanReq) (*yunweiclient.OpenPlanCommonResp, error) {

	err := l.svcCtx.OpenPlanModel.TransactInsert(l.ctx, in.OpenPlatData)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}
	return &yunweiclient.OpenPlanCommonResp{}, nil
}
