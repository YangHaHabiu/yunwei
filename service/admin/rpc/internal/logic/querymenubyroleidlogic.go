package logic

import (
	"context"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuByRoleIdLogic {
	return &QueryMenuByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(in *adminclient.QueryMenuByRoleIdReq) (*adminclient.QueryMenuByRoleIdResp, error) {
	RoleMenus, _ := l.svcCtx.RoleMenuModel.FindByRoleId(l.ctx, in.Id)
	var list []int64
	for _, user := range *RoleMenus {
		list = append(list, user.MenuId)
	}
	return &adminclient.QueryMenuByRoleIdResp{
		Ids: list,
	}, nil
}
