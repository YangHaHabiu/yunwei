package user

import (
	"context"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/api/internal/logic/common"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAssignmentPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAssignmentPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAssignmentPolicyLogic {
	return &UserAssignmentPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAssignmentPolicyLogic) UserAssignmentPolicy(req *types.UserAssignmentPolicyReq) error {
	_, err := l.svcCtx.AdminRpc.UserAssignmentPolicy(l.ctx, &adminclient.UserAssignmentPolicyReq{
		UserId:    req.UserId,
		UserCheck: req.UserChecked,
	})
	if err != nil {
		return err
	}
	//刷新策略
	common.FlushStrategy(l.svcCtx, l.ctx, ctxdata.GetUidFromCtx(l.ctx), ctxdata.GetUnameFromCtx(l.ctx))

	return nil
}
