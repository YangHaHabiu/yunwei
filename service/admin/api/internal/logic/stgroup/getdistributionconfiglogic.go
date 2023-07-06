package stgroup

import (
	"context"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDistributionConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDistributionConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDistributionConfigLogic {
	return &GetDistributionConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDistributionConfigLogic) GetDistributionConfig(req *types.GetDistributionConfigReq) (resp *types.GetDistributionConfigResp, err error) {

	info, err := l.svcCtx.AdminRpc.GetUserCheckStategyInfo(l.ctx, &adminclient.StgroupUserCheckInfoReq{
		Id: req.StgroupId,
	})
	if err != nil {
		return nil, err
	}
	tmpuser := make([]int64, 0)
	for _, v := range info.UserCheck {
		tmpuser = append(tmpuser, v.Id)
	}
	tmpugroup := make([]int64, 0)
	for _, v := range info.UgroupCheck {
		tmpugroup = append(tmpugroup, v.Id)
	}
	listUser, err := l.svcCtx.AdminRpc.UserList(l.ctx, &adminclient.UserListReq{})
	if err != nil {
		return nil, err
	}
	tmpa := make([]*types.NewUserData, 0)
	for _, v := range listUser.List {
		tmpa = append(tmpa, &types.NewUserData{
			Id:   v.Id,
			Name: v.NickName + "(" + v.Name + ")",
		})
	}
	listUgroup, err := l.svcCtx.AdminRpc.UgroupList(l.ctx, &adminclient.UgroupListReq{})
	if err != nil {
		return nil, err
	}
	tmpg := make([]*types.NewUserData, 0)
	for _, v := range listUgroup.List {
		tmpg = append(tmpg, &types.NewUserData{
			Id:   v.Id,
			Name: v.UgName,
		})
	}

	return &types.GetDistributionConfigResp{
		UserChecked:      tmpuser,
		UserGroupChecked: tmpugroup,
		UserAllData:      tmpa,
		UserGroupAllData: tmpg,
	}, nil
}
