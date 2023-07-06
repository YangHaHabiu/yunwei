package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// loginlog rpc start
func (l *LoginLogAddLogic) LoginLogAdd(in *adminclient.LoginLogAddReq) (*adminclient.LoginLogAddResp, error) {
	_, err := l.svcCtx.LoginLogModel.Insert(l.ctx, &model.SysLoginLog{
		UserName: in.UserName,
		Status:   in.Status,
		Ip:       in.Ip,
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
	}

	return &adminclient.LoginLogAddResp{}, nil
}
