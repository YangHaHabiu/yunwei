package user

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePersonalDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePersonalDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonalDataLogic {
	return &UpdatePersonalDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePersonalDataLogic) UpdatePersonalData(req *types.UpdatePersonalReq) (err error) {
	_, err = l.svcCtx.AdminRpc.UpdatePersonalInfo(l.ctx, &admin.UserUpdatePersonalInfoReq{
		Id:       ctxdata.GetUidFromCtx(l.ctx),
		Email:    req.Email,
		Avatar:   req.Avatar,
		NickName: req.NickName,
		Mobile:   req.Mobile,
	})
	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("修改用户信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}

	return
}
