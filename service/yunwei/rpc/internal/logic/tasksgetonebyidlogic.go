package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksGetOneByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksGetOneByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksGetOneByIdLogic {
	return &TasksGetOneByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TasksGetOneByIdLogic) TasksGetOneById(in *yunweiclient.GetTasksReq) (*yunweiclient.ListTasksData, error) {

	one, err := l.svcCtx.TasksModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条任务失败")
	}

	var tmp yunweiclient.ListTasksData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("复制单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
