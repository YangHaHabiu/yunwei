package insideProxyHost

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProxyHostDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideProxyHostDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProxyHostDeleteLogic {
	return &InsideProxyHostDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideProxyHostDeleteLogic) InsideProxyHostDelete(req *types.DeleteInsideProxyHostReq) error {
	_, err := l.svcCtx.IntranetRpc.InsideProxyHostDelete(l.ctx, &intranetclient.DeleteInsideProxyHostReq{InsideProxyHostId: req.InsideProxyHostId})
	if err != nil {
		return err
	}
	return nil
}
