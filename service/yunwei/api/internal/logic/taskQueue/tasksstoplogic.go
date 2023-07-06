package taskQueue

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksStopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksStopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksStopLogic {
	return &TasksStopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksStopLogic) TasksStop(req *types.GetTasksReq) error {
	_, err := l.svcCtx.YunWeiRpc.TaskStop(l.ctx, &yunweiclient.StopTasksReq{Id: req.Id})
	if err != nil {
		return err
	}
	return nil
}
