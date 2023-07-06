package asset

import (
	"context"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OwnerProjectDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOwnerProjectDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OwnerProjectDataLogic {
	return &OwnerProjectDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OwnerProjectDataLogic) OwnerProjectData() (*types.OwnerProjectDataResp, error) {

	projectIdsList := make([]*types.OwnerProjectData, 0)
	//根据用户id查询对应的项目
	ownerList, err := l.svcCtx.AdminRpc.ProjectOwnerList(l.ctx, &adminclient.ProjectOwnerReq{UserId: ctxdata.GetUidFromCtx(l.ctx)})
	if err != nil {
		return nil, err
	}

	if len(ownerList.List) != 0 {
		for _, v := range ownerList.List {
			projectIdsList = append(projectIdsList, &types.OwnerProjectData{
				Label:     v.ProjectCn,
				Value:     v.ProjectId,
				ProjectEn: v.ProjectEn,
			})
		}
	}

	return &types.OwnerProjectDataResp{OwnerProjectData: projectIdsList}, nil
}
