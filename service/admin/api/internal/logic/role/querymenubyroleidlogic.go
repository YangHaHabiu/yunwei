package role

import (
	"context"
	"github.com/jinzhu/copier"
	"strconv"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuByRoleIdLogic {
	return &QueryMenuByRoleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(req *types.RoleMenuReq) (*types.RoleMenuResp, error) {
	//查询所有菜单
	resp, _ := l.svcCtx.AdminRpc.MenuList(l.ctx, &admin.MenuListReq{
		Name: "",
		Url:  "",
	})

	var (
		list    []*types.ListMenuData
		listIds []int64
	)
	listIds = make([]int64, 0)
	for _, menu := range resp.List {
		list = append(list, &types.ListMenuData{
			Key:      strconv.FormatInt(menu.Id, 10),
			Title:    menu.Name,
			ParentId: menu.ParentId,
			Id:       menu.Id,
			Label:    menu.Name,
		})
	}

	QueryMenu, _ := l.svcCtx.AdminRpc.QueryMenuByRoleId(l.ctx, &admin.QueryMenuByRoleIdReq{
		Id: req.RoleId,
	})

	copier.Copy(&listIds, QueryMenu.Ids)

	return &types.RoleMenuResp{
		AllData:  list,
		RoleData: listIds,
	}, nil
}
