package ugroup

import (
	"context"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UgroupAssignmentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUgroupAssignmentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupAssignmentUserLogic {
	return &UgroupAssignmentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UgroupAssignmentUserLogic) UgroupAssignmentUser(req *types.UgroupAssignmentUserReq) error {
	_, err := l.svcCtx.AdminRpc.UgroupAssignmentUser(l.ctx, &adminclient.UgroupAssignmentUserReq{
		UserCheck: req.UserChecked,
		UgroupId:  req.UgroupId,
	})

	return err
}
