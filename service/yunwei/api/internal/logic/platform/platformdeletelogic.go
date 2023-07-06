package platform

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlatformDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlatformDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformDeleteLogic {
	return &PlatformDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlatformDeleteLogic) PlatformDelete(req *types.DeletePlatformReq) error {
	_, err := l.svcCtx.YunWeiRpc.PlatformDelete(l.ctx, &yunweiclient.DeletePlatformReq{PlatformId: req.PlatformId})
	if err != nil {
		return err
	}
	return nil
}
