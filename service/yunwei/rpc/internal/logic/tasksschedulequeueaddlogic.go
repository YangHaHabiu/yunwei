package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksScheduleQueueAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksScheduleQueueAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueAddLogic {
	return &TasksScheduleQueueAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TasksScheduleQueue Rpc Start
func (l *TasksScheduleQueueAddLogic) TasksScheduleQueueAdd(in *yunweiclient.AddTasksScheduleQueueReq) (*yunweiclient.TasksScheduleQueueCommonResp, error) {
	var tmp model.TasksScheduleQueue
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	err = l.svcCtx.TasksScheduleQueueModel.Insert(l.ctx, &tmp, in.Maps)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}

	return &yunweiclient.TasksScheduleQueueCommonResp{}, nil
}
