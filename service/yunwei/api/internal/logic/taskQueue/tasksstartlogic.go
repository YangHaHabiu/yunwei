package taskQueue

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksStartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksStartLogic {
	return &TasksStartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksStartLogic) TasksStart(req *types.GetTasksReq) error {
	_, err := l.svcCtx.YunWeiRpc.TaskStart(l.ctx, &yunweiclient.StartTasksReq{Id: req.Id})
	if err != nil {
		return err
	}
	return nil
}
