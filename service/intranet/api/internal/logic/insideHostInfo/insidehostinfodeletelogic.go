package insideHostInfo

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideHostInfoDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideHostInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideHostInfoDeleteLogic {
	return &InsideHostInfoDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideHostInfoDeleteLogic) InsideHostInfoDelete(req *types.DeleteInsideHostInfoReq) error {
	_, err := l.svcCtx.IntranetRpc.InsideHostInfoDelete(l.ctx, &intranetclient.DeleteInsideHostInfoReq{InsideHostInfoId: req.InsideHostInfoId})
	if err != nil {
		return err
	}
	return nil
}
