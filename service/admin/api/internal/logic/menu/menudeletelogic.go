package menu

import (
	"context"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDeleteLogic) MenuDelete(req *types.DeleteMenuReq) (err error) {
	_, err = l.svcCtx.AdminRpc.MenuDelete(l.ctx, &admin.MenuDeleteReq{
		Id: req.MenuId,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据menuId: %d,删除菜单异常:%s", req.MenuId, err.Error())
		return err
	}

	return nil
}
