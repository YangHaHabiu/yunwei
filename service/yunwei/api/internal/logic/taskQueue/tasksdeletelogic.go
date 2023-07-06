package taskQueue

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksDeleteLogic {
	return &TasksDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksDeleteLogic) TasksDelete(req *types.DeleteTasksReq) error {
	_, err := l.svcCtx.YunWeiRpc.TasksDelete(l.ctx, &yunweiclient.DeleteTasksReq{Id: req.Id})
	if err != nil {
		return err
	}
	return nil
}
