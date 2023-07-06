package user

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePersonalPasswordDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePersonalPasswordDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonalPasswordDataLogic {
	return &UpdatePersonalPasswordDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePersonalPasswordDataLogic) UpdatePersonalPasswordData(req *types.UpdatePersonalPasswordReq) (err error) {

	if req.NewPasswordRepeat != req.NewPassword {
		return xerr.NewErrMsg("两次输入的密码不一致，请检查")
	}
	if !tool.VerifyPassword(8, 18, req.NewPasswordRepeat) {
		return xerr.NewErrMsg("不符合密码校验规则: 必须包含数字、大写字母、小写字母、特殊字符(如.@$!%*#_~?&^)至少3种的组合且长度在8-18之间")
	}

	_, err = l.svcCtx.AdminRpc.UpdatePersonalPassword(l.ctx, &admin.UserUpdatePersonalPasswordReq{
		Id:                ctxdata.GetUidFromCtx(l.ctx),
		OldPassword:       req.OldPassword,
		NewPasswordRepeat: req.NewPasswordRepeat,
		NewPassword:       req.NewPassword,
	})
	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("修改用户密码失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return
}
