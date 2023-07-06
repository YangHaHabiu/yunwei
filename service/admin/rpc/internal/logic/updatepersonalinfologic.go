package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePersonalInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePersonalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonalInfoLogic {
	return &UpdatePersonalInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePersonalInfoLogic) UpdatePersonalInfo(in *adminclient.UserUpdatePersonalInfoReq) (*adminclient.UserUpdatePersonalInfoResp, error) {
	err := l.svcCtx.UserModel.UpdatePersonalInfo(l.ctx, &model.SysUser{
		Id:       in.Id,
		Email:    in.Email,
		Avatar:   in.Avatar,
		NickName: in.NickName,
		Mobile:   in.Mobile,
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.ADMIN_UPDATEPERSON_ERROR)
	}
	return &adminclient.UserUpdatePersonalInfoResp{}, nil
}
