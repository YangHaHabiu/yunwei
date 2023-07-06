package tasksScheduleQueue

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksScheduleQueueAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksScheduleQueueAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueAddLogic {
	return &TasksScheduleQueueAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksScheduleQueueAddLogic) TasksScheduleQueueAdd(req *types.AddTasksScheduleQueueReq) error {
	var tmp yunweiclient.TasksScheduleQueueCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	err, b := common.GetScheduleType(l.svcCtx, l.ctx)
	if err != nil {
		return err
	}

	_, err = l.svcCtx.YunWeiRpc.TasksScheduleQueueAdd(l.ctx, &yunweiclient.AddTasksScheduleQueueReq{One: &tmp, Maps: b})
	if err != nil {
		return err
	}
	return nil
}
