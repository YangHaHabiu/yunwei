package taskQueue

import (
	"context"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WaitMergeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWaitMergeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WaitMergeListLogic {
	return &WaitMergeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WaitMergeListLogic) WaitMergeList(req *types.ListWaitMergeReq) (resp *types.ListWaitMergeResp, err error) {

	return
}
