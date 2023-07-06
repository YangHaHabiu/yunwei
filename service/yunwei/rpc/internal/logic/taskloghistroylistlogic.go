package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskLogHistroyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskLogHistroyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskLogHistroyListLogic {
	return &TaskLogHistroyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Tasks Rpc End
func (l *TaskLogHistroyListLogic) TaskLogHistroyList(in *yunweiclient.ListTaskLogHistroyReq) (*yunweiclient.ListTaskLogHistroyResp, error) {

	filters := make([]interface{}, 0)
	filters = append(filters,
		"pid__=", in.TaskId,
	)
	all, err := l.svcCtx.TaskLogHistroyModel.FindAll(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询任务历史日志失败，原因：" + err.Error())
	}

	if len(*all) != 1 {
		l.Logger.Error(all)
		return nil, xerr.NewErrMsg("任务历史日志失败")
	}

	return &yunweiclient.ListTaskLogHistroyResp{
		Data: (*all)[0].LogInfo,
	}, nil
}
