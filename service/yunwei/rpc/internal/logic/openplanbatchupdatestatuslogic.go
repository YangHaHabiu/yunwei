package logic

import (
	"context"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenplanBatchUpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenplanBatchUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenplanBatchUpdateStatusLogic {
	return &OpenplanBatchUpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenplanBatchUpdateStatusLogic) OpenplanBatchUpdateStatus(in *yunweiclient.BatchUpdateStatusReq) (*yunweiclient.BatchUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &yunweiclient.BatchUpdateStatusResp{}, nil
}
