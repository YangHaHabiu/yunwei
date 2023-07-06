package insideVersion

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideVersionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideVersionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideVersionDeleteLogic {
	return &InsideVersionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideVersionDeleteLogic) InsideVersionDelete(req *types.DeleteInsideVersionReq) error {
	_, err := l.svcCtx.IntranetRpc.InsideVersionDelete(l.ctx, &intranetclient.DeleteInsideVersionReq{InsideVersionId: req.InsideVersionId})
	if err != nil {
		return err
	}
	return nil
}
