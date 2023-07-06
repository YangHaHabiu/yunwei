package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskLogHistroyDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskLogHistroyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskLogHistroyDetailLogic {
	return &TaskLogHistroyDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskLogHistroyDetailLogic) TaskLogHistroyDetail(in *yunweiclient.DetailTaskLogHistroyReq) (*yunweiclient.DetailTaskLogHistroyResp, error) {
	one, err := l.svcCtx.TaskLogHistroyModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条日志记录失败" + err.Error())
	}

	var tmp yunweiclient.DetailTaskLogHistroyResp
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("复制单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
