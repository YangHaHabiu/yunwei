package asset

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetDeleteLogic {
	return &AssetDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetDeleteLogic) AssetDelete(req *types.DeleteAssetReq) error {
	_, err := l.svcCtx.YunWeiRpc.AssetDelete(l.ctx, &yunweiclient.AssetDeleteReq{
		AssetId: req.AssetId,
	})
	if err != nil {
		return err
	}
	return nil
}
