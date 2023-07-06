package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUgroupAssignmentUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUgroupAssignmentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUgroupAssignmentUserLogic {
	return &GetUgroupAssignmentUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUgroupAssignmentUserLogic) GetUgroupAssignmentUser(in *adminclient.GetUgroupAssignmentUserReq) (*adminclient.GetUgroupAssignmentUserResp, error) {

	all, err := l.svcCtx.UserUgroupModel.FindAll(l.ctx, "ugroup_id__=", in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询用户组关联用户失败" + err.Error())
	}
	tmp := make([]*adminclient.UgroupAssignmentUserData, 0)
	for _, v := range *all {
		tmp = append(tmp, &adminclient.UgroupAssignmentUserData{
			UgroupId: v.UgroupId,
			UserId:   v.UserId,
		})
	}

	return &adminclient.GetUgroupAssignmentUserResp{
		Data: tmp,
	}, nil
}
