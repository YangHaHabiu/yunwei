package role

import (
	"context"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req *types.DeleteRoleReq) error {
	_, err := l.svcCtx.AdminRpc.RoleDelete(l.ctx, &admin.RoleDeleteReq{
		Id: req.RoleId,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据roleId: %d,删除角色异常:%s", req.RoleId, err.Error())
		return err
	}

	return nil
}
