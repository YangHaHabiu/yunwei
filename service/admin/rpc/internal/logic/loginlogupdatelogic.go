package logic

import (
	"context"
	"fmt"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogUpdateLogic {
	return &LoginLogUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogUpdateLogic) LoginLogUpdate(in *adminclient.LoginLogUpdateReq) (*adminclient.LoginLogUpdateResp, error) {

	err := l.svcCtx.LoginLogModel.UpdateFieldStatusByUname(l.ctx, in.UserName, in.Status)
	if err != nil {
		return nil, xerr.NewErrMsg(fmt.Sprintf("修改登录日志失败，原因：%v", err.Error()))
	}

	return &adminclient.LoginLogUpdateResp{}, nil
}
