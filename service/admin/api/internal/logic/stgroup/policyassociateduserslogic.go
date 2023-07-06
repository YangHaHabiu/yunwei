package stgroup

import (
	"context"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/api/internal/logic/common"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
	"ywadmin-v3/service/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type PolicyAssociatedUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPolicyAssociatedUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PolicyAssociatedUsersLogic {
	return &PolicyAssociatedUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PolicyAssociatedUsersLogic) PolicyAssociatedUsers(req *types.PolicyAssociatedUsersReq) (err error) {
	_, err = l.svcCtx.AdminRpc.PolicyAssociatedUsers(l.ctx, &admin.PolicyAssociatedUsersReq{
		UserCheck:   req.UserChecked,
		UgroupCheck: req.UserGroupChecked,
		StgroupId:   req.StgroupId,
	})
	if err != nil {
		return err
	}
	//刷新策略
	common.FlushStrategy(l.svcCtx, l.ctx, ctxdata.GetUidFromCtx(l.ctx), ctxdata.GetUnameFromCtx(l.ctx))
	return
}
