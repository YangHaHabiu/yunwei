package menu

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUpdateLogic) MenuUpdate(req *types.UpdateMenuReq) error {
	_, err := l.svcCtx.AdminRpc.MenuUpdate(l.ctx, &admin.MenuUpdateReq{
		Id:           req.Id,
		Name:         req.Name,
		ParentId:     req.ParentId,
		Url:          req.Url,
		Perms:        req.Perms,
		Type:         req.Type,
		Icon:         req.Icon,
		OrderNum:     req.OrderNum,
		LastUpdateBy: ctxdata.GetUnameFromCtx(l.ctx),
		VuePath:      req.VuePath,
		VueComponent: req.VueComponent,
		VueIcon:      req.VueIcon,
		VueRedirect:  req.VueRedirect,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新菜单信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return nil

}
