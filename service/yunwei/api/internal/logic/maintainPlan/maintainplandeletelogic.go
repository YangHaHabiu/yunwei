package maintainPlan

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaintainPlanDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanDeleteLogic {
	return &MaintainPlanDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaintainPlanDeleteLogic) MaintainPlanDelete(req *types.DeleteMaintainPlanReq) error {

	_, err := l.svcCtx.YunWeiRpc.MaintainPlanDelete(l.ctx, &yunweiclient.DeleteMaintainPlanReq{Id: req.Id})
	if err != nil {
		return err
	}
	return nil
}
