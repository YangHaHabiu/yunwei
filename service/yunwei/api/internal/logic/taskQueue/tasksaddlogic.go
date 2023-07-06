package taskQueue

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksAddLogic {
	return &TasksAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksAddLogic) TasksAdd(req *types.AddTasksReq) error {
	var tmp yunweiclient.AddTasksReq
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝任务数据失败，原因：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.TasksAdd(l.ctx, &tmp)
	if err != nil {
		return err
	}
	return nil
}
