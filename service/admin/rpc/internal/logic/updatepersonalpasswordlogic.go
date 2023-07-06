package logic

import (
	"context"
	"fmt"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePersonalPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePersonalPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonalPasswordLogic {
	return &UpdatePersonalPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePersonalPasswordLogic) UpdatePersonalPassword(in *adminclient.UserUpdatePersonalPasswordReq) (*adminclient.UserUpdatePersonalPasswordResp, error) {

	//校验旧密码
	userInfo, err2 := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err2 != nil {
		return nil, xerr.NewErrMsg("查不到此用户的信息，请检查")
	}
	inputPwd := tool.Md5ByString(in.OldPassword + userInfo.Salt + userInfo.Name)
	fmt.Println(inputPwd)
	if inputPwd != userInfo.Password {
		return nil, xerr.NewErrMsg("旧密码输入错误，请检查")
	}
	//确认两次新密码一致
	if in.NewPassword != in.NewPasswordRepeat {
		return nil, xerr.NewErrMsg("两次输入密码有误，请检查")
	}

	//修改新密码
	err := l.svcCtx.UserModel.UpdatePassword(l.ctx, &model.SysUser{
		Id:       in.Id,
		Password: tool.Md5ByString(in.NewPassword + userInfo.Salt + userInfo.Name),
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ADMIN_UPDATEPERSON_PASSWORD_ERROR)
	}

	return &adminclient.UserUpdatePersonalPasswordResp{}, nil
}
