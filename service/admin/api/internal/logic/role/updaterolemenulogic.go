package role

import (
	"context"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenuLogic {
	return &UpdateRoleMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleMenuLogic) UpdateRoleMenu(req *types.UpdateRoleMenuReq) error {
	_, err := l.svcCtx.AdminRpc.UpdateMenuRole(l.ctx, &admin.UpdateMenuRoleReq{
		RoleId:  req.RoleId,
		MenuIds: req.MenuIds,
	})

	if err != nil {
		return err
	}

	return nil
}
