package logic

import (
	"context"
	"time"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuUpdateLogic) MenuUpdate(in *adminclient.MenuUpdateReq) (*adminclient.MenuUpdateResp, error) {
	err := l.svcCtx.MenuModel.Update(l.ctx, &model.SysMenu{
		Id:             in.Id,
		Name:           in.Name,
		ParentId:       in.ParentId,
		Url:            in.Url,
		Perms:          in.Perms,
		Tp:             in.Type,
		Icon:           in.Icon,
		OrderNum:       in.OrderNum,
		CreateTime:     time.Time{},
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
		VuePath:        in.VuePath,
		VueComponent:   in.VueComponent,
		VueIcon:        in.VueIcon,
		VueRedirect:    in.VueRedirect,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return &adminclient.MenuUpdateResp{}, nil
}
