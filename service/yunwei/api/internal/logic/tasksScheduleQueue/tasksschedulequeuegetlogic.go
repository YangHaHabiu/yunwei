package tasksScheduleQueue

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksScheduleQueueGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksScheduleQueueGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueGetLogic {
	return &TasksScheduleQueueGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksScheduleQueueGetLogic) TasksScheduleQueueGet(req *types.GetTasksScheduleQueueReq) (resp *types.ListTasksScheduleQueueData, err error) {
	get, err := l.svcCtx.YunWeiRpc.TasksScheduleQueueGet(l.ctx, &yunweiclient.GetTasksScheduleQueueReq{TasksScheduleQueueId: req.TasksScheduleQueueId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListTasksScheduleQueueData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
