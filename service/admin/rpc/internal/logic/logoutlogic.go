package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/identity/rpc/identity"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *adminclient.LogoutReq) (*adminclient.LogoutResp, error) {
	_, err := l.svcCtx.IdentityRpc.ClearToken(l.ctx, &identity.ClearTokenReq{
		UserId: in.UserId,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("用户退出失败")
	}
	return &adminclient.LogoutResp{}, nil
}
