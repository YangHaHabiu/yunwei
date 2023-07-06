package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"
	"ywadmin-v3/service/identity/rpc/identity"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// user rpc start
func (l *LoginLogic) Login(in *adminclient.LoginReq) (*adminclient.LoginResp, error) {

	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Username)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, xerr.NewErrCode(xerr.ADMIN_USERNAME_ERROR)
	default:
		return nil, xerr.NewErrCode(xerr.ADMIN_USERNAME_ERROR)
	}

	inputPwd := tool.Md5ByString(in.Password + userInfo.Salt + userInfo.Name)
	if userInfo.Password != inputPwd {
		return nil, xerr.NewErrCode(xerr.ADMIN_USERNAMEPWD_ERROR)
	}
	resp, err := l.svcCtx.IdentityRpc.GenerateToken(l.ctx, &identity.GenerateTokenReq{
		UserId:   userInfo.Id,
		UserName: userInfo.Name,
		NickName: userInfo.NickName,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR)
	}

	return &adminclient.LoginResp{
		Username:     userInfo.Name,
		Id:           userInfo.Id,
		AccessToken:  resp.AccessToken,
		AccessExpire: resp.AccessExpire,
		RefreshAfter: resp.RefreshAfter,
	}, nil
}
