package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksScheduleQueueGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksScheduleQueueGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueGetLogic {
	return &TasksScheduleQueueGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TasksScheduleQueueGetLogic) TasksScheduleQueueGet(in *yunweiclient.GetTasksScheduleQueueReq) (*yunweiclient.ListTasksScheduleQueueData, error) {
	one, err := l.svcCtx.TasksScheduleQueueModel.FindOne(l.ctx, in.TasksScheduleQueueId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp yunweiclient.ListTasksScheduleQueueData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
