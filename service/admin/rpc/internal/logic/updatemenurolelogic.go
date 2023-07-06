package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuRoleLogic {
	return &UpdateMenuRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuRoleLogic) UpdateMenuRole(in *adminclient.UpdateMenuRoleReq) (*adminclient.UpdateMenuRoleResp, error) {
	err := l.svcCtx.RoleMenuModel.Delete(l.ctx, in.RoleId)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_DELETE_ERROR)
	}
	for _, id := range in.MenuIds {
		_, err = l.svcCtx.RoleMenuModel.Insert(l.ctx, &model.SysRoleMenu{
			RoleId: in.RoleId,
			MenuId: id,
		})
		if err != nil {
			return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
		}
	}

	return &adminclient.UpdateMenuRoleResp{}, nil
}
