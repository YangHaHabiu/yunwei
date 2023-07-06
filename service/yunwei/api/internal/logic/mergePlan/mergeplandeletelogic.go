package mergePlan

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergePlanDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergePlanDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanDeleteLogic {
	return &MergePlanDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergePlanDeleteLogic) MergePlanDelete(req *types.DeleteMergePlanReq) error {
	_, err := l.svcCtx.YunWeiRpc.MergePlanDelete(l.ctx, &yunweiclient.DeleteMergePlanReq{Id: req.Id})
	if err != nil {
		return err
	}
	return nil
}
