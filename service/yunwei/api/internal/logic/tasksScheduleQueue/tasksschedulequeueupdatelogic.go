package tasksScheduleQueue

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunwei"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksScheduleQueueUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksScheduleQueueUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueUpdateLogic {
	return &TasksScheduleQueueUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksScheduleQueueUpdateLogic) TasksScheduleQueueUpdate(req *types.UpdateTasksScheduleQueueReq) error {
	var tmp yunwei.TasksScheduleQueueCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("更新拷贝数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.TasksScheduleQueueUpdate(l.ctx, &yunweiclient.UpdateTasksScheduleQueueReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
