package insideOperation

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideOperationDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideOperationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideOperationDeleteLogic {
	return &InsideOperationDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideOperationDeleteLogic) InsideOperationDelete(req *types.DeleteInsideOperationReq) error {
	_, err := l.svcCtx.IntranetRpc.InsideOperationDelete(l.ctx, &intranetclient.DeleteInsideOperationReq{InsideOperationId: req.InsideOperationId})
	if err != nil {
		return err
	}
	return nil
}
