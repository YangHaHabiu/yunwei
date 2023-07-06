package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *adminclient.InfoReq) (*adminclient.InfoResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, xerr.NewErrCode(xerr.ADMIN_NOTFOUNDUID_ERROR)
	default:
		return nil, err
	}
	var list []*adminclient.MenuListTree
	if userInfo.Name == globalkey.SuperUserName {

		menus, err := l.svcCtx.MenuModel.FindAll(l.ctx, "tp__in", "0,1")
		if err != nil {
			return nil, xerr.NewErrCode(xerr.ADMIN_MENUSELECT_ERROR)
		}

		list = listTrees(menus, list)
	} else {
		menus, err := l.svcCtx.MenuModel.FindAllByUserId(l.ctx, in.UserId)
		if err != nil {
			return nil, xerr.NewErrCode(xerr.ADMIN_MENUSELECT_ERROR)
		}
		list = listTrees(menus, list)
	}
	return &adminclient.InfoResp{
		Avatar:       userInfo.Avatar,
		Name:         userInfo.NickName,
		MenuListTree: list,
	}, nil
}

func listTrees(menus *[]model.SysMenu, list []*adminclient.MenuListTree) []*adminclient.MenuListTree {
	for _, menu := range *menus {
		list = append(list, &adminclient.MenuListTree{
			Id:           menu.Id,
			Name:         menu.Name,
			Icon:         menu.Icon,
			ParentId:     menu.ParentId,
			Path:         menu.Url,
			VuePath:      menu.VuePath,
			VueComponent: menu.VueComponent,
			VueIcon:      menu.VueIcon,
			OrderNum:     menu.OrderNum,
			VueRedirect:  menu.VueRedirect,
			TableName:    menu.TableName,
			IsShow:       menu.IsShow,
			KeepAlive:    menu.KeepAlive,
		})
	}
	return list
}
