package insideServer

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideServerDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideServerDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideServerDeleteLogic {
	return &InsideServerDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideServerDeleteLogic) InsideServerDelete(req *types.DeleteInsideServerReq) error {
	_, err := l.svcCtx.IntranetRpc.InsideServerDelete(l.ctx, &intranetclient.DeleteInsideServerReq{InsideServerId: req.InsideServerId})
	if err != nil {
		return err
	}
	return nil
}
