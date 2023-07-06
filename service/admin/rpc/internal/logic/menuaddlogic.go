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

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// user rpc end
func (l *MenuAddLogic) MenuAdd(in *adminclient.MenuAddReq) (*adminclient.MenuAddResp, error) {
	_, err := l.svcCtx.MenuModel.Insert(l.ctx, &model.SysMenu{
		Id:             0,
		Name:           in.Name,
		ParentId:       in.ParentId,
		Url:            in.Url,
		Perms:          in.Perms,
		Tp:             in.Type,
		Icon:           in.Icon,
		OrderNum:       in.OrderNum,
		CreateBy:       in.CreateBy,
		CreateTime:     time.Time{},
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
		VuePath:        in.VuePath,
		VueComponent:   in.VueComponent,
		VueIcon:        in.VueIcon,
		VueRedirect:    in.VueRedirect,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
	}

	return &adminclient.MenuAddResp{}, nil
}
