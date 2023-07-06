package logic

import (
	"context"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReSetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReSetPasswordLogic) ReSetPassword(in *adminclient.ReSetPasswordReq) (*adminclient.ReSetPasswordResp, error) {
	one, err2 := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	if err2 != nil || one == nil {
		return nil, xerr.NewErrCode(xerr.ADMIN_USERNAME_ERROR)
	}
	err := l.svcCtx.UserModel.UpdatePassword(l.ctx, &model.SysUser{
		Id:           in.Id,
		Password:     tool.Md5ByString(in.NewPassword + one.Salt + one.Name),
		LastUpdateBy: in.LastUpdateBy,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}
	return &adminclient.ReSetPasswordResp{}, nil
}
