package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuListLogic) MenuList(in *adminclient.MenuListReq) (*adminclient.MenuListResp, error) {
	count, _ := l.svcCtx.MenuModel.Count(l.ctx)
	all, err := l.svcCtx.MenuModel.FindAll(l.ctx,
		"name__like", in.Name, "url__like", in.Url,
	)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询菜单列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrCode(xerr.ADMIN_MENUSELECT_ERROR)
	}
	var list []*adminclient.MenuListData
	for _, menu := range *all {
		list = append(list, &adminclient.MenuListData{
			Id:             menu.Id,
			Name:           menu.Name,
			ParentId:       menu.ParentId,
			Url:            menu.Url,
			Perms:          menu.Perms,
			Type:           menu.Tp,
			Icon:           menu.Icon,
			OrderNum:       menu.OrderNum,
			CreateBy:       menu.CreateBy,
			CreateTime:     menu.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   menu.LastUpdateBy,
			LastUpdateTime: menu.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        menu.DelFlag,
			VuePath:        menu.VuePath,
			VueComponent:   menu.VueComponent,
			VueIcon:        menu.VueIcon,
			VueRedirect:    menu.VueRedirect,
			TableName:      menu.TableName,
			IsShow:         menu.IsShow,
			KeepAlive:      menu.KeepAlive,
		})
	}

	return &adminclient.MenuListResp{
		Total: count,
		List:  list,
	}, nil
}
