package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUgroupAssignmentPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUgroupAssignmentPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUgroupAssignmentPolicyLogic {
	return &GetUgroupAssignmentPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUgroupAssignmentPolicyLogic) GetUgroupAssignmentPolicy(in *adminclient.GetUgroupAssignmentPolicyReq) (*adminclient.GetUgroupAssignmentPolicyResp, error) {

	all, err := l.svcCtx.StgroupUgroupModel.FindAll(l.ctx, "ugroup_id__=", in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询用户组关联策略失败" + err.Error())
	}
	tmp := make([]*adminclient.UgroupAssignmentPolicyData, 0)
	for _, v := range *all {
		tmp = append(tmp, &adminclient.UgroupAssignmentPolicyData{
			UgroupId:  v.UgroupId,
			StgroupId: v.StgroupId,
		})
	}
	return &adminclient.GetUgroupAssignmentPolicyResp{
		Data: tmp,
	}, nil
}
