package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksScheduleQueueDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksScheduleQueueDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueDeleteLogic {
	return &TasksScheduleQueueDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TasksScheduleQueueDeleteLogic) TasksScheduleQueueDelete(in *yunweiclient.DeleteTasksScheduleQueueReq) (*yunweiclient.TasksScheduleQueueCommonResp, error) {
	err := l.svcCtx.TasksScheduleQueueModel.DeleteSoft(l.ctx, in.TasksScheduleQueueId, in.Maps)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}

	return &yunweiclient.TasksScheduleQueueCommonResp{}, nil
}
