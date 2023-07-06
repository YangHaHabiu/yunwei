package menu

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuAddLogic) MenuAdd(req *types.AddMenuReq) (err error) {
	_, err = l.svcCtx.AdminRpc.MenuAdd(l.ctx, &admin.MenuAddReq{
		Name:         req.Name,
		ParentId:     req.ParentId,
		Url:          req.Url,
		Perms:        req.Perms,
		Type:         req.Type,
		Icon:         req.Icon,
		OrderNum:     req.OrderNum,
		CreateBy:     "admin",
		VuePath:      req.VuePath,
		VueComponent: req.VueComponent,
		VueIcon:      req.VueIcon,
		VueRedirect:  req.VueRedirect,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加菜单信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return nil
}
