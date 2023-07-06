package user

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAssignmentPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserAssignmentPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAssignmentPolicyLogic {
	return &GetUserAssignmentPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAssignmentPolicyLogic) GetUserAssignmentPolicy(req *types.GetUserAssignmentPolicyReq) (resp *types.GetUserAssignmentPolicyResp, err error) {
	policy, err := l.svcCtx.AdminRpc.GetUserAssignmentPolicy(l.ctx, &adminclient.GetUserAssignmentPolicyReq{Id: req.UserId})
	if err != nil {
		return nil, err
	}
	tmp := make([]int64, 0)
	for _, v := range policy.Data {
		tmp = append(tmp, v.StgroupId)
	}
	list, err := l.svcCtx.AdminRpc.StgroupList(l.ctx, &adminclient.StgroupListReq{})
	if err != nil {
		return nil, err
	}
	var tmps []*types.ListStgroupData
	copier.Copy(&tmps, list.List)

	return &types.GetUserAssignmentPolicyResp{
		UserChecked: tmp,
		UserAllData: tmps,
	}, nil
}
