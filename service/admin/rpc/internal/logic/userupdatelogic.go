package logic

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *adminclient.UserUpdateReq) (*adminclient.UserUpdateResp, error) {
	_ = l.svcCtx.UserModel.Update(l.ctx, &model.SysUser{
		Id:           in.Id,
		NickName:     in.NickName,
		Avatar:       in.Avatar,
		Email:        in.Email,
		Mobile:       in.Mobile,
		DeptId:       in.DeptId,
		Status:       in.Status,
		LastUpdateBy: in.LastUpdateBy,
	})

	roleObj := strings.Split(in.RoleIds, ",")
	userRoleList := make([]*model.SysUserRole, 0)
	if len(in.RoleIds) != 0 {
		for _, v := range roleObj {
			if gconv.Int64(v) != 0 {
				userRoleList = append(userRoleList, &model.SysUserRole{
					UserId: in.Id,
					RoleId: gconv.Int64(v),
				})
			}

		}
		_ = l.svcCtx.UserRoleModel.DeleteByCustomFiled(l.ctx, []interface{}{
			"user_id__=",
			in.Id,
		}...)
		_ = l.svcCtx.UserRoleModel.BulkInserter(userRoleList)
	} else {
		_ = l.svcCtx.UserRoleModel.DeleteByCustomFiled(l.ctx, []interface{}{
			"user_id__=",
			in.Id,
		}...)
	}

	ugroupObj := strings.Split(in.UgroupIds, ",")
	userUgroupList := make([]*model.SysUserUgroup, 0)
	if len(in.UgroupIds) != 0 {
		for _, v := range ugroupObj {
			if gconv.Int64(v) != 0 {
				userUgroupList = append(userUgroupList, &model.SysUserUgroup{
					UserId:   in.Id,
					UgroupId: gconv.Int64(v),
				})
			}
		}

		_ = l.svcCtx.UserUgroupModel.DeleteByCustomFiled(l.ctx, []interface{}{
			"user_id__=",
			in.Id,
		}...)
		_ = l.svcCtx.UserUgroupModel.BulkInserter(userUgroupList)
	} else {
		_ = l.svcCtx.UserUgroupModel.DeleteByCustomFiled(l.ctx, []interface{}{
			"user_id__=",
			in.Id,
		}...)
	}
	projectObj := strings.Split(in.ProjectIds, ",")
	userProjectList := make([]*model.SysUserProject, 0)
	if len(in.ProjectIds) != 0 {
		for _, v := range projectObj {
			if gconv.Int64(v) != 0 {
				userProjectList = append(userProjectList, &model.SysUserProject{
					UserId:    in.Id,
					ProjectId: gconv.Int64(v),
				})
			}
		}
		_ = l.svcCtx.UserProjectModel.DeleteByCustomFiled(l.ctx, []interface{}{
			"user_id__=",
			in.Id,
		}...)
		_ = l.svcCtx.UserProjectModel.BulkInserter(userProjectList)
	} else {
		_ = l.svcCtx.UserProjectModel.DeleteByCustomFiled(l.ctx, []interface{}{
			"user_id__=",
			in.Id,
		}...)
	}
	return &adminclient.UserUpdateResp{}, nil
}
