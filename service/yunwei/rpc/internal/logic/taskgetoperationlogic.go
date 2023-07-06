package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskGetOperationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskGetOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskGetOperationLogic {
	return &TaskGetOperationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskGetOperationLogic) TaskGetOperation(in *yunweiclient.TaskGetOperationReq) (*yunweiclient.TaskGetOperationResp, error) {

	list, err := l.svcCtx.TasksModel.GetOperationList(l.ctx, in.Uid)
	if err != nil {
		return nil, xerr.NewErrMsg("查询操作权限失败，原因：" + err.Error())
	}

	var (
		tmp []*yunweiclient.TaskGetOperationData
	)
	err = copier.Copy(&tmp, list)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	return &yunweiclient.TaskGetOperationResp{
		List: tmp,
	}, nil
}
