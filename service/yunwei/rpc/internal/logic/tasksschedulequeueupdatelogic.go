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

type TasksScheduleQueueUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksScheduleQueueUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueUpdateLogic {
	return &TasksScheduleQueueUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TasksScheduleQueueUpdateLogic) TasksScheduleQueueUpdate(in *yunweiclient.UpdateTasksScheduleQueueReq) (*yunweiclient.TasksScheduleQueueCommonResp, error) {
	var tmp model.TasksScheduleQueue
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝更新数据失败，原因：" + err.Error())
	}
	err = l.svcCtx.TasksScheduleQueueModel.Update(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}

	return &yunweiclient.TasksScheduleQueueCommonResp{}, nil
}
