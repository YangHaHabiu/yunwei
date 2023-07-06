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

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UpdateUserReq) error {
	_, err := l.svcCtx.AdminRpc.UserUpdate(l.ctx, &admin.UserUpdateReq{
		Id:           req.Id,
		Email:        req.Email,
		Mobile:       req.Mobile,
		Name:         req.Name,
		NickName:     req.NickName,
		DeptId:       req.DeptId,
		LastUpdateBy: ctxdata.GetUnameFromCtx(l.ctx),
		RoleIds:      req.RoleIds,
		UgroupIds:    req.UgroupIds,
		ProjectIds:   req.ProjectIds,
		Status:       req.Status,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新用户信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}

	return nil
}
