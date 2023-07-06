package featureServer

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeatureServerDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeatureServerDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeatureServerDeleteLogic {
	return &FeatureServerDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeatureServerDeleteLogic) FeatureServerDelete(req *types.DeleteFeatureServerReq) error {
	_, err := l.svcCtx.YunWeiRpc.FeatureServerDelete(l.ctx, &yunweiclient.DeleteFeatureServerReq{FeatureServerId: req.FeatureServerId})
	if err != nil {
		return err
	}

	return nil
}
