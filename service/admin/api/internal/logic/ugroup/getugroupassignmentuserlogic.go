package ugroup

import (
	"context"
	"fmt"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUgroupAssignmentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUgroupAssignmentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUgroupAssignmentUserLogic {
	return &GetUgroupAssignmentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUgroupAssignmentUserLogic) GetUgroupAssignmentUser(req *types.GetUgroupAssignmentUserReq) (resp *types.GetUgroupAssignmentUserResp, err error) {
	policy, err := l.svcCtx.AdminRpc.GetUgroupAssignmentUser(l.ctx, &adminclient.GetUgroupAssignmentUserReq{Id: req.UgroupId})
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
	tmps := make([]*types.ListUserDataNew, 0)
	for _, v := range list.List {
		tmps = append(tmps, &types.ListUserDataNew{
			Id:     v.Id,
			StName: fmt.Sprintf("%s(%s)", v.NickName, v.Name),
		})
	}

	return &types.GetUgroupAssignmentUserResp{
		UserChecked: tmp,
		UserAllData: tmps,
	}, nil
}
