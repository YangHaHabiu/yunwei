package user

import (
	"context"
	"net/http"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/myip"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"github.com/zeromicro/go-zero/core/logx"
	"ywadmin-v3/service/admin/api/internal/svc"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(r *http.Request) (err error) {
	_, err = l.svcCtx.AdminRpc.Logout(l.ctx, &adminclient.LogoutReq{
		UserId: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return err
	}
	//退出系统日志记录
	l.svcCtx.AdminRpc.LoginLogAdd(l.ctx, &adminclient.LoginLogAddReq{
		UserName: ctxdata.GetUnameFromCtx(l.ctx),
		Status:   "logout",
		Ip:       myip.GetCurrentIP(r),
	})
	//并修改登录日志
	l.svcCtx.AdminRpc.LoginLogUpdate(l.ctx, &adminclient.LoginLogUpdateReq{
		UserName: ctxdata.GetUnameFromCtx(l.ctx),
		Status:   "login",
	})
	return
}
