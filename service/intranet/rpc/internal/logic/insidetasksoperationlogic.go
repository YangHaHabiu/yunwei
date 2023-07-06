package logic

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksOperationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksOperationLogic {
	return &InsideTasksOperationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksOperationLogic) InsideTasksOperation(in *intranetclient.InsideTasksOperationReq) (*intranetclient.InsideTasksOperationResp, error) {
	var msg string
	if in.OperationType == "getlog" {
		one, err := l.svcCtx.InsideTasksLogsModel.FindOne(l.ctx, in.TasksId)
		if err != nil {
			return nil, xerr.NewErrMsg("查询任务日志失败")
		}
		msg = one.Content
	} else if in.OperationType == "stop" {
		one, err := l.svcCtx.InsideTasksPidModel.FindOne(l.ctx, in.TasksId)
		if err != nil {
			return nil, xerr.NewErrMsg("停止失败")
		}
		//杀掉进程
		msg = gconv.String(one.TasksPid)
	}

	return &intranetclient.InsideTasksOperationResp{
		Pong: msg,
	}, nil
}
