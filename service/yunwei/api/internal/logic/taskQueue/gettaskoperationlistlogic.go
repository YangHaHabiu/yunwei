package taskQueue

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskOperationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskOperationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskOperationListLogic {
	return &GetTaskOperationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskOperationListLogic) GetTaskOperationList() (resp *types.GetTaskOperationList, err error) {

	list, err := l.svcCtx.YunWeiRpc.TaskGetOperation(l.ctx, &yunweiclient.TaskGetOperationReq{Uid: ctxdata.GetUidFromCtx(l.ctx)})
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.TaskOperationListData, 0)

	err = copier.Copy(&tmp, list.List)
	if err != nil {
		return nil, err
	}

	resp = new(types.GetTaskOperationList)
	resp.OpTreeData = tmp
	return

}
