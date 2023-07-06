package asset

import (
	"context"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetListLogic {
	return &AssetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetListLogic) AssetList(req *types.ListAssetReq) (*types.ListAssetResp, error) {
	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds, req.ProjectType)
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.ListAssetData, 0)
	var total int64

	list, err := l.svcCtx.YunWeiRpc.AssetList(l.ctx, &yunweiclient.AssetListReq{
		Current:             req.Current,
		PageSize:            req.PageSize,
		RecycleType:         req.RecycleType,
		AssetId:             req.AssetId,
		Ips:                 req.Ips,
		ProjectIds:          projectIds,
		OwnershipCompanyIds: req.OwnershipCompanyIds,
		InitType:            req.InitType,
		CleanType:           req.CleanType,
		Provider:            req.Provider,
		Label:               req.Label,
		HostRoleCn:          req.HostRoleCn,
	})

	if err != nil {
		return nil, err
	}

	err = copier.Copy(&tmp, list.List)
	if err != nil {
		return nil, err
	}
	total = list.Total

	//出机方
	companyList, err := common.GetCompanyList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}
	//云商
	providerList, err := common.GetProviderList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}
	//用途
	hostRoleList, err := common.GetHostRoleList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}

	//自定义筛选条件
	filterList := []*types.FilterList{

		{
			Label:    "出机方",
			Value:    "ownershipCompanyIds",
			Types:    "select",
			Children: companyList,
		},
		{
			Label:    "项目",
			Value:    "projectIds",
			Types:    "select",
			Children: projectList,
		},
		{
			Label: "IP",
			Value: "ips",
			Types: "input",
		},
		{
			Label:    "云商",
			Value:    "provider",
			Types:    "select",
			Children: providerList,
		},
		{
			Label:    "用途",
			Value:    "hostRoleCn",
			Types:    "select",
			Children: hostRoleList,
		},
		{
			Label: "标签",
			Value: "label",
			Types: "input",
		},
		{
			Label: "初始化状态",
			Value: "initType",
			Types: "select",
			Children: []*types.FilterList{
				{
					Label: "已初始化",
					Value: "1",
				},
				{
					Label: "未初始化",
					Value: "2",
				},
			},
		},
		{
			Label: "清理状态",
			Value: "cleanType",
			Types: "select",
			Children: []*types.FilterList{
				{
					Label: "已清理",
					Value: "1",
				},
				{
					Label: "未清理",
					Value: "2",
				},
			},
		},
	}

	return &types.ListAssetResp{
		Rows:   tmp,
		Total:  total,
		Filter: filterList,
	}, nil
}
