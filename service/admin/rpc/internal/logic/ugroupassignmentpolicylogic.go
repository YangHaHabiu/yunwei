package logic

import (
	"context"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UgroupAssignmentPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUgroupAssignmentPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupAssignmentPolicyLogic {
	return &UgroupAssignmentPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UgroupAssignmentPolicyLogic) UgroupAssignmentPolicy(in *adminclient.UgroupAssignmentPolicyReq) (*adminclient.UgroupAssignmentPolicyResp, error) {
	err := l.svcCtx.StgroupUgroupModel.TransactInsert(l.ctx, &adminclient.PolicyAssociatedUsersReq{
		UgroupCheck: in.UgroupCheck,
		StgroupId:   in.UgroupId,
	}, "opts")
	if err != nil {
		return nil, err
	}
	return &adminclient.UgroupAssignmentPolicyResp{}, nil
}
