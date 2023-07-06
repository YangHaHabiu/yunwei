package insideInstallPlan

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideInstallPlanDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideInstallPlanDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideInstallPlanDeleteLogic {
	return &InsideInstallPlanDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideInstallPlanDeleteLogic) InsideInstallPlanDelete(req *types.DeleteInsideInstallPlanReq) error {
	_, err := l.svcCtx.IntranetRpc.InsideInstallPlanDelete(l.ctx, &intranetclient.DeleteInsideInstallPlanReq{InsideInstallPlanId: req.InsideInstallPlanId})
	if err != nil {
		return err
	}
	return nil
}
