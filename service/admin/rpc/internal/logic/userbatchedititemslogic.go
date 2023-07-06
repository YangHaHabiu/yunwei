package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserBatchEditItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserBatchEditItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserBatchEditItemsLogic {
	return &UserBatchEditItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserBatchEditItemsLogic) UserBatchEditItems(in *adminclient.UserBatchEditItemsReq) (*adminclient.UserBatchEditItemsResp, error) {

	err := l.svcCtx.UserModel.UserBatchEditItems(l.ctx, in)
	if err != nil {
		return nil, xerr.NewErrMsg("批量操作用户与项目表失败，原因:" + err.Error())
	}
	return &adminclient.UserBatchEditItemsResp{}, nil
}
