package asset

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunwei"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetRecycleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetRecycleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetRecycleDeleteLogic {
	return &AssetRecycleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetRecycleDeleteLogic) AssetRecycleDelete(req *types.RecycleDeleteAssetReq) error {
	_, err := l.svcCtx.YunWeiRpc.AssetRecycleDelete(l.ctx, &yunwei.AssetRecycleDeleteReq{
		AssetId:     req.AssetId,
		RecycleType: 1,
	})
	if err != nil {
		return err
	}

	return nil
}
