package ugroup

import (
	"context"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UgroupAssignmentPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUgroupAssignmentPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupAssignmentPolicyLogic {
	return &UgroupAssignmentPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UgroupAssignmentPolicyLogic) UgroupAssignmentPolicy(req *types.UgroupAssignmentPolicyReq) error {
	_, err := l.svcCtx.AdminRpc.UgroupAssignmentPolicy(l.ctx, &adminclient.UgroupAssignmentPolicyReq{
		UgroupCheck: req.UgroupChecked,
		UgroupId:    req.UgroupId,
	})
	if err != nil {
		return err
	}

	return nil
}
