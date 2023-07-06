package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleAssignmentUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleAssignmentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleAssignmentUserLogic {
	return &GetRoleAssignmentUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleAssignmentUserLogic) GetRoleAssignmentUser(in *adminclient.GetRoleAssignmentUserReq) (*adminclient.GetRoleAssignmentUserResp, error) {
	all, err := l.svcCtx.UserRoleModel.FindAll(l.ctx, "role_id__=", in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询角色关联用户失败" + err.Error())
	}
	tmp := make([]*adminclient.RoleAssignmentUserData, 0)
	for _, v := range *all {
		tmp = append(tmp, &adminclient.RoleAssignmentUserData{
			RoleId: v.RoleId,
			UserId: v.UserId,
		})
	}
	return &adminclient.GetRoleAssignmentUserResp{
		Data: tmp,
	}, nil
}
