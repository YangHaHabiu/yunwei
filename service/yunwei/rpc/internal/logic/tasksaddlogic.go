package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksAddLogic {
	return &TasksAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ServerName Rpc End
func (l *TasksAddLogic) TasksAdd(in *yunweiclient.AddTasksReq) (*yunweiclient.TasksCommonResp, error) {
	//事务插入数据
	err := l.svcCtx.TasksModel.Insert(l.ctx, in, l.svcCtx.Config.YwQQGroup)
	if err != nil {
		return nil, xerr.NewErrMsg("新增任务失败，原因：" + err.Error())
	}

	return &yunweiclient.TasksCommonResp{}, nil
}
