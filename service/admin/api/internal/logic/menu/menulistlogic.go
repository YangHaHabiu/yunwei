package menu

import (
	"context"
	"encoding/json"
	"strconv"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListLogic) MenuList(req *types.ListMenuReq) (*types.ListMenuResp, error) {
	resp, err := l.svcCtx.AdminRpc.MenuList(l.ctx, &admin.MenuListReq{
		Name: req.Name,
		Url:  req.Url,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询菜单列表异常:%s", string(data), err.Error())
		return nil, err
	}

	list := make([]*types.ListtMenuData, 0)

	for _, menu := range resp.List {
		list = append(list, &types.ListtMenuData{
			Id:             menu.Id,
			Key:            strconv.FormatInt(menu.Id, 10),
			Name:           menu.Name,
			Title:          menu.Name,
			ParentId:       menu.ParentId,
			Url:            menu.Url,
			Perms:          menu.Perms,
			Type:           menu.Type,
			Icon:           menu.Icon,
			OrderNum:       menu.OrderNum,
			CreateBy:       menu.CreateBy,
			CreateTime:     menu.CreateTime,
			LastUpdateBy:   menu.LastUpdateBy,
			LastUpdateTime: menu.LastUpdateTime,
			VuePath:        menu.VuePath,
			VueComponent:   menu.VueComponent,
			VueIcon:        menu.VueIcon,
			VueRedirect:    menu.VueRedirect,
			TableName:      menu.TableName,
			IsShow:         menu.IsShow,
			KeepAlive:      menu.KeepAlive,
		})
	}

	return &types.ListMenuResp{
		Rows:  list,
		Total: resp.Total,
	}, nil
}
