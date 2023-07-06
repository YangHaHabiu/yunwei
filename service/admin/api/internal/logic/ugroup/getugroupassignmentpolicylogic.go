package ugroup

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUgroupAssignmentPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUgroupAssignmentPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUgroupAssignmentPolicyLogic {
	return &GetUgroupAssignmentPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUgroupAssignmentPolicyLogic) GetUgroupAssignmentPolicy(req *types.GetUgroupAssignmentPolicyReq) (resp *types.GetUgroupAssignmentPolicyResp, err error) {
	policy, err := l.svcCtx.AdminRpc.GetUgroupAssignmentPolicy(l.ctx, &adminclient.GetUgroupAssignmentPolicyReq{Id: req.UgroupId})
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

	return &types.GetUgroupAssignmentPolicyResp{
		UgroupChecked: tmp,
		UgroupAllData: tmps,
	}, nil
}
