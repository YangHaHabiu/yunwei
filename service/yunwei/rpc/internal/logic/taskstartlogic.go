package logic

import (
	"context"
	"time"
	"ywadmin-v3/service/yunwei/rpc/jobs"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskStartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskStartLogic {
	return &TaskStartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskStartLogic) TaskStart(in *yunweiclient.StartTasksReq) (*yunweiclient.StartTasksResp, error) {
	err := TaskCommon(l.svcCtx, l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	ch := make(chan string)
	ctx, _ := context.WithCancel(context.Background())
	go jobs.Run(l.svcCtx, ctx, in.Id, 10*time.Minute, ch)
	return &yunweiclient.StartTasksResp{}, nil
}
