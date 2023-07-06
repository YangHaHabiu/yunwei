package logic

import (
	"context"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAssignmentPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAssignmentPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAssignmentPolicyLogic {
	return &UserAssignmentPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAssignmentPolicyLogic) UserAssignmentPolicy(in *adminclient.UserAssignmentPolicyReq) (*adminclient.UserAssignmentPolicyResp, error) {
	err := l.svcCtx.StgroupUserModel.TransactInsert(l.ctx, &adminclient.PolicyAssociatedUsersReq{
		UserCheck: in.UserCheck,
		StgroupId: in.UserId,
	}, "opts")
	if err != nil {
		return nil, err
	}

	return &adminclient.UserAssignmentPolicyResp{}, nil
}
