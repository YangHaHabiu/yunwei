package logic

import (
	"context"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PolicyAssociatedUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPolicyAssociatedUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PolicyAssociatedUsersLogic {
	return &PolicyAssociatedUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PolicyAssociatedUsersLogic) PolicyAssociatedUsers(in *adminclient.PolicyAssociatedUsersReq) (*adminclient.PolicyAssociatedUsersResp, error) {
	err := l.svcCtx.StgroupUserModel.TransactInsert(l.ctx, in)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.StgroupUgroupModel.TransactInsert(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return &adminclient.PolicyAssociatedUsersResp{}, nil
}
