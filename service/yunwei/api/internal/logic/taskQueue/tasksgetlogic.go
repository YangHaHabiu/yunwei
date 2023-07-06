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

type TasksGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksGetLogic {
	return &TasksGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksGetLogic) TasksGet(req *types.GetTasksReq) (resp *types.GetTasksResp, err error) {
	list, err := l.svcCtx.YunWeiRpc.TasksGet(l.ctx, &yunweiclient.GetTasksReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListTasksData, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	errList := make([]int64, 0)
	allList := make([]int64, 0)
	copier.Copy(&errList, list.ErrArr)
	copier.Copy(&allList, list.AllIdArr)
	resp = new(types.GetTasksResp)
	resp.Rows = tmp
	resp.ErrArr = errList
	resp.AllIdArr = allList

	return
}
