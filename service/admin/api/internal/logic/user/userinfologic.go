package user

import (
	"context"
	"strings"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/api/internal/logic/common"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {

	// 这里的key和生成jwt token时传入的key一致
	userId := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.AdminRpc.UserInfo(l.ctx, &admin.InfoReq{
		UserId: userId,
	})

	if err != nil {
		return nil, err
	}
	list, _ := l.svcCtx.AdminRpc.UserList(l.ctx, &adminclient.UserListReq{UserId: userId, Current: 1, PageSize: 1})
	//组装element ui中的菜单
	MenuTreeVue := make([]*types.ListMenuTreeVue, 0)
	for _, item := range resp.MenuListTree {
		if len(strings.TrimSpace(item.VuePath)) != 0 {
			MenuTreeVue = append(MenuTreeVue, &types.ListMenuTreeVue{
				Id:           item.Id,
				ParentId:     item.ParentId,
				Title:        item.Name,
				Path:         item.VuePath,
				Name:         item.Name,
				Icon:         item.VueIcon,
				VueRedirect:  item.VueRedirect,
				VueComponent: item.VueComponent,
				IsShow:       item.IsShow,
				Meta: types.MenuTreeMeta{
					Title:     item.Name,
					Icon:      item.VueIcon,
					OrderNum:  item.OrderNum,
					KeepAlive: item.KeepAlive,
				},
			})
		}
	}
	//刷新策略
	common.FlushStrategy(l.svcCtx, l.ctx, ctxdata.GetUidFromCtx(l.ctx), ctxdata.GetUnameFromCtx(l.ctx))

	return &types.UserInfoResp{
		WebTitle:    l.svcCtx.Config.VerificationCodeWatermark,
		Avatar:      resp.Avatar,
		Name:        resp.Name,
		UserId:      userId,
		MenuTreeVue: MenuTreeVue,
		Profiles: types.ListUserData{
			Id:          list.List[0].Id,
			Name:        list.List[0].Name,
			NickName:    list.List[0].NickName,
			Avatar:      list.List[0].NickName,
			Email:       list.List[0].Email,
			Mobile:      list.List[0].Mobile,
			DeptId:      list.List[0].DeptId,
			CreateBy:    list.List[0].CreateBy,
			CreateTime:  list.List[0].CreateTime,
			RoleName:    list.List[0].RoleName,
			DeptName:    list.List[0].DeptName,
			UgroupIds:   list.List[0].UgroupIds,
			UgroupNames: list.List[0].UgroupNames,
		},
	}, nil
}
