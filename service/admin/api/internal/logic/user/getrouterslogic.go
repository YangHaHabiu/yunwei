package user

import (
	"context"
	"sort"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/xsort"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoutersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoutersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoutersLogic {
	return &GetRoutersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoutersLogic) GetRouters() (resp *types.GetRoutersResp, err error) {
	info, err := l.svcCtx.AdminRpc.UserInfo(l.ctx, &adminclient.InfoReq{
		UserId: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	return getMenu(info)
}

func getMenu(resp *adminclient.InfoResp) (*types.GetRoutersResp, error) {
	//一级菜单
	oneMenus := make([]*adminclient.MenuListTree, 0)
	for _, v := range resp.MenuListTree {
		if v.ParentId == 0 {
			oneMenus = append(oneMenus, v)
		}
	}
	//根据order num从小到大排序
	sort.Sort(xsort.OneList(oneMenus))
	lists := make([]*types.MenuNewVue, 0)
	for _, v1 := range oneMenus {
		//二级菜单
		twoMenus := make([]*adminclient.MenuListTree, 0)
		for _, v2 := range resp.MenuListTree {
			if v2.ParentId == v1.Id {
				twoMenus = append(twoMenus, v2)
			}
		}
		sort.Sort(xsort.OneList(twoMenus))
		childrenList := make([]*types.MenuNewVue, 0)
		for _, v := range twoMenus {
			childrenList = append(childrenList, &types.MenuNewVue{
				Path:      v.VuePath,
				Name:      v.VuePath,
				Hidden:    false,
				Component: v.VueComponent,
				Meta: types.MetaVue{
					Icon:  v.VueIcon,
					Title: v.Name,
				},
			})
		}
		lists = append(lists, &types.MenuNewVue{
			AlwaysShow: false,
			Component:  v1.VueComponent,
			Hidden:     false,
			Meta: types.MetaVue{
				Icon:  v1.VueIcon,
				Title: v1.Name,
			},
			Children: childrenList,
			Path:     v1.Path,
		})
	}

	return &types.GetRoutersResp{
		Rows: lists,
	}, nil
}
