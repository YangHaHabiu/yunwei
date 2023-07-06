package tasksScheduleQueue

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksScheduleQueueDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksScheduleQueueDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueDeleteLogic {
	return &TasksScheduleQueueDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksScheduleQueueDeleteLogic) TasksScheduleQueueDelete(req *types.DeleteTasksScheduleQueueReq) error {

	err, b := common.GetScheduleType(l.svcCtx, l.ctx)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.YunWeiRpc.TasksScheduleQueueDelete(l.ctx, &yunweiclient.DeleteTasksScheduleQueueReq{
		TasksScheduleQueueId: req.TasksScheduleQueueId,
		Maps:                 b,
	})
	if err != nil {
		return err
	}
	return nil
}
