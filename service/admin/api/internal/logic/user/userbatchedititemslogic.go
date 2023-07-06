package user

import (
	"context"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserBatchEditItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserBatchEditItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserBatchEditItemsLogic {
	return &UserBatchEditItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserBatchEditItemsLogic) UserBatchEditItems(req *types.UserBatchEditItemsReq) error {
	_, err := l.svcCtx.AdminRpc.UserBatchEditItems(l.ctx, &adminclient.UserBatchEditItemsReq{
		ProjectIds: req.ProjectIds,
		Operate:    req.Operate,
		UserIds:    req.UserIds,
	})

	return err
}
