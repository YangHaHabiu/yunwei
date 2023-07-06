package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *adminclient.UserDeleteReq) (*adminclient.UserDeleteResp, error) {

	//删除用户时，先删除用户相关的所有关联
	err := l.svcCtx.UserModel.TransactDelete(l.ctx, &model.SysUser{
		Id:           in.Id,
		LastUpdateBy: in.LastUpdateBy,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("删除用户及相关的关联表数据失败，原因：" + err.Error())
	}

	return &adminclient.UserDeleteResp{}, nil
}
