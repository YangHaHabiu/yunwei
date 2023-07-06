package insideProjectCluster

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProjectClusterDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideProjectClusterDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProjectClusterDeleteLogic {
	return &InsideProjectClusterDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideProjectClusterDeleteLogic) InsideProjectClusterDelete(req *types.DeleteInsideProjectClusterReq) error {
	_, err := l.svcCtx.IntranetRpc.InsideProjectClusterDelete(l.ctx, &intranetclient.DeleteInsideProjectClusterReq{InsideProjectClusterId: req.InsideProjectClusterId})
	if err != nil {
		return err
	}
	return nil
}
