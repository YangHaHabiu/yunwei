package insideTasksLogs

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksLogsDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideTasksLogsDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksLogsDeleteLogic {
	return &InsideTasksLogsDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideTasksLogsDeleteLogic) InsideTasksLogsDelete(req *types.DeleteInsideTasksLogsReq) error {
	_, err := l.svcCtx.IntranetRpc.InsideTasksLogsDelete(l.ctx, &intranetclient.DeleteInsideTasksLogsReq{InsideTasksLogsId: req.InsideTasksLogsId})
	if err != nil {
		return err
	}
	return nil
}
