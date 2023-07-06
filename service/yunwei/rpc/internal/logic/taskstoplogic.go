package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/jobs"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskStopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskStopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskStopLogic {
	return &TaskStopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskStopLogic) TaskStop(in *yunweiclient.StopTasksReq) (*yunweiclient.StopTasksResp, error) {
	err := TaskCommon(l.svcCtx, l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithCancel(context.Background())
	err = jobs.Stop(l.svcCtx, ctx, in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("停止任务失败，原因：" + err.Error())
	}
	return &yunweiclient.StopTasksResp{}, nil
}
