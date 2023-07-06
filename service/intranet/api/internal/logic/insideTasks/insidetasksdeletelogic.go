package insideTasks

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideTasksDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksDeleteLogic {
	return &InsideTasksDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideTasksDeleteLogic) InsideTasksDelete(req *types.DeleteInsideTasksReq) error {
	_, err := l.svcCtx.IntranetRpc.InsideTasksDelete(l.ctx, &intranetclient.DeleteInsideTasksReq{InsideTasksId: req.InsideTasksId})
	if err != nil {
		return err
	}
	return nil
}
