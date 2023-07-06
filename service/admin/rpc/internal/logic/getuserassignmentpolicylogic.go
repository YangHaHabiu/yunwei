package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAssignmentPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAssignmentPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAssignmentPolicyLogic {
	return &GetUserAssignmentPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAssignmentPolicyLogic) GetUserAssignmentPolicy(in *adminclient.GetUserAssignmentPolicyReq) (*adminclient.GetUserAssignmentPolicyResp, error) {

	all, err := l.svcCtx.StgroupUserModel.FindAll(l.ctx, "user_id__=", in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询用户关联策略失败" + err.Error())
	}
	tmp := make([]*adminclient.UserAssignmentPolicyData, 0)
	for _, v := range *all {
		tmp = append(tmp, &adminclient.UserAssignmentPolicyData{
			UserId:    v.UserId,
			StgroupId: v.StgroupId,
		})
	}
	return &adminclient.GetUserAssignmentPolicyResp{
		Data: tmp,
	}, nil
}
