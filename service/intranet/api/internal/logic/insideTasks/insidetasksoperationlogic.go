package insideTasks

import (
	"context"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideTasksOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksOperationLogic {
	return &InsideTasksOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideTasksOperationLogic) InsideTasksOperation(req *types.OperationInsideTasksReq) (resp *types.OperationInsideTasksResp, err error) {
	operation, err := l.svcCtx.IntranetRpc.InsideTasksOperation(l.ctx, &intranetclient.InsideTasksOperationReq{
		TasksId:       req.TasksId,
		OperationType: req.OperationType,
	})
	if err != nil {
		return nil, err
	}
	return &types.OperationInsideTasksResp{
		Row: operation.Pong,
	}, nil
}
