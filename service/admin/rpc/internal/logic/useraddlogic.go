package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *adminclient.UserAddReq) (*adminclient.UserAddResp, error) {
	name, err := l.svcCtx.UserModel.FindOneOnlyWithName(l.ctx, in.Name)

	if name != nil {
		return nil, xerr.NewErrMsg("存在相同的用户名[" + in.Name + "]，请检查")
	}
	err = l.svcCtx.UserModel.TransactInsert(l.ctx, in)
	if err != nil {
		return nil, xerr.NewErrMsg("插入用户数据失败，原因：" + err.Error())
	}
	return &adminclient.UserAddResp{}, nil
}
