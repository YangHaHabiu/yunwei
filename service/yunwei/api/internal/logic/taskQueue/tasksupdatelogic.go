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

type TasksUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksUpdateLogic {
	return &TasksUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksUpdateLogic) TasksUpdate(req *types.UpdateTasksReq) error {
	var tmp yunweiclient.UpdateTasksReq
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("更新拷贝任务数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.TasksUpdate(l.ctx, &tmp)
	if err != nil {
		return err
	}
	return nil
}
