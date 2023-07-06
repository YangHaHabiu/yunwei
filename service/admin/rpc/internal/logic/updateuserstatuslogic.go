package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserStatusLogic) UpdateUserStatus(in *adminclient.UserStatusReq) (*adminclient.UserStatusResp, error) {
	err := l.svcCtx.UserModel.UpdateStatus(l.ctx, &model.SysUser{
		Id:           in.Id,
		Status:       in.Status,
		LastUpdateBy: in.LastUpdateBy,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("修改用户状态失败")
	}
	return &adminclient.UserStatusResp{}, nil
}
