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

type TaskGetFormatJsonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskGetFormatJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskGetFormatJsonLogic {
	return &TaskGetFormatJsonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskGetFormatJsonLogic) TaskGetFormatJson(req *types.TaskGetFormatJsonReq) (resp *types.TaskGetFormatJsonResp, err error) {

	var tmp yunweiclient.TaskGetFormatJsonReq
	err = copier.Copy(&tmp, req)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝格式化任务数据失败，原因：" + err.Error())
	}
	json, err := l.svcCtx.YunWeiRpc.TaskGetFormatJson(l.ctx, &tmp)
	if err != nil {
		return nil, err
	}

	tmpResp := make([]*types.OperationListM, 0)
	err = copier.Copy(&tmpResp, json.OperationListM)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝格式化任务数据失败，原因：" + err.Error())
	}

	resp = new(types.TaskGetFormatJsonResp)
	resp.OperationListM = tmpResp

	return
}
