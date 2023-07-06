package openPlan

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenPlanDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenPlanDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanDeleteLogic {
	return &OpenPlanDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenPlanDeleteLogic) OpenPlanDelete(req *types.DeleteOpenPlanReq) error {
	_, err := l.svcCtx.YunWeiRpc.OpenPlanDelete(l.ctx, &yunweiclient.DeleteOpenPlanReq{Id: req.Id})
	if err != nil {
		return err
	}
	return nil
}
