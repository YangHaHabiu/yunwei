package role

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleAssignmentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleAssignmentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleAssignmentUserLogic {
	return &GetRoleAssignmentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleAssignmentUserLogic) GetRoleAssignmentUser(req *types.GetRoleAssignmentUserReq) (resp *types.GetRoleAssignmentUserResp, err error) {
	policy, err := l.svcCtx.AdminRpc.GetRoleAssignmentUser(l.ctx, &adminclient.GetRoleAssignmentUserReq{Id: req.RoleId})
	if err != nil {
		return nil, err
	}
	tmp := make([]int64, 0)
	for _, v := range policy.Data {
		tmp = append(tmp, v.UserId)
	}
	list, err := l.svcCtx.AdminRpc.UserList(l.ctx, &adminclient.UserListReq{
		Current: 0, PageSize: 0,
	})
	if err != nil {
		return nil, err
	}
	var tmps []*types.ListUserData
	copier.Copy(&tmps, list.List)

	return &types.GetRoleAssignmentUserResp{
		UserChecked: tmp,
		UserAllData: tmps,
	}, nil
}
