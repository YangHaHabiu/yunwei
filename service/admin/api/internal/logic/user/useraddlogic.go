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

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.AddUserReq) (err error) {
	_, err = l.svcCtx.AdminRpc.UserAdd(l.ctx, &admin.UserAddReq{
		Email:      req.Email,
		Mobile:     req.Mobile,
		Name:       req.Name,
		NickName:   req.NickName,
		DeptId:     req.DeptId,
		CreateBy:   ctxdata.GetUnameFromCtx(l.ctx),
		RoleIds:    req.RoleIds,
		UgroupIds:  req.UgroupIds,
		ProjectIds: req.ProjectIds,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加用户信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}

	return nil
}
