package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAssignmentUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAssignmentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAssignmentUserLogic {
	return &RoleAssignmentUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAssignmentUserLogic) RoleAssignmentUser(in *adminclient.RoleAssignmentUserReq) (*adminclient.RoleAssignmentUserResp, error) {

	err := l.svcCtx.UserRoleModel.TransactInsert(l.ctx, in, "opts")
	if err != nil {
		return nil, xerr.NewErrMsg("更新角色与用户关联数据失败，原因：" + err.Error())
	}
	return &adminclient.RoleAssignmentUserResp{}, nil
}
