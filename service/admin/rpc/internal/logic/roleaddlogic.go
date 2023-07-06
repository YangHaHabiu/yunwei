package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// role rpc start
func (l *RoleAddLogic) RoleAdd(in *adminclient.RoleAddReq) (*adminclient.RoleAddResp, error) {
	_, err := l.svcCtx.RoleModel.Insert(l.ctx, &model.SysRole{
		Id:           0,
		Name:         in.Name,
		Remark:       in.Remark,
		CreateBy:     in.CreateBy,
		LastUpdateBy: in.CreateBy,
		DelFlag:      0,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
	}

	return &adminclient.RoleAddResp{}, nil
}
