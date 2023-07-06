package role

import (
	"context"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAssignmentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAssignmentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAssignmentUserLogic {
	return &RoleAssignmentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAssignmentUserLogic) RoleAssignmentUser(req *types.RoleAssignmentUserReq) error {
	_, err := l.svcCtx.AdminRpc.RoleAssignmentUser(l.ctx, &adminclient.RoleAssignmentUserReq{
		UserCheck: req.UserChecked,
		RoleId:    req.RoleId,
	})

	return err
}
