package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UgroupAssignmentUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUgroupAssignmentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupAssignmentUserLogic {
	return &UgroupAssignmentUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UgroupAssignmentUserLogic) UgroupAssignmentUser(in *adminclient.UgroupAssignmentUserReq) (*adminclient.UgroupAssignmentUserResp, error) {

	err := l.svcCtx.UserUgroupModel.TransactInsert(l.ctx, in, "opts")
	if err != nil {
		return nil, xerr.NewErrMsg("更新用户组关联用户失败，原因：" + err.Error())
	}
	return &adminclient.UgroupAssignmentUserResp{}, nil
}
