package openPlan

import (
	"context"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenPlanListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenPlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanListLogic {
	return &OpenPlanListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenPlanListLogic) OpenPlanList(req *types.ListOpenPlanReq) (resp *types.ListOpenPlanResp, err error) {
	sortTmpList := make([]*yunweiclient.SortFiledData, 0)
	err = copier.Copy(&sortTmpList, req.SortFields)
	if err != nil {
		return nil, err
	}

	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListOpenPlanData, 0)

	list, err := l.svcCtx.YunWeiRpc.OpenPlanList(l.ctx, &yunweiclient.ListOpenPlanReq{
		Current:       req.Current,
		PageSize:      req.PageSize,
		DateRange:     req.DateRange,
		ProjectIds:    projectIds,
		PlatformIds:   req.PlatformIds,
		InitdbStatus:  req.InitdbStatus,
		InstallStatus: req.InstallStatus,
		ClusterName:   req.ClusterName,
		SortFiledList: sortTmpList,
		GameType:      req.GameType,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}
	//获取平台
	platformList, err := common.GetPlatform(l.svcCtx, l.ctx, projectIds, "", req.PlatformType)
	if err != nil {
		return nil, err
	}
	//获取集群
	clusterList, err := common.GetCluster(l.svcCtx, l.ctx, projectIds)
	if err != nil {
		return nil, err
	}

	fl, err := common.GetInstallStatusList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}
	fl2, err := common.GetInitdbStatusList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}

	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label:    "项目",
			Value:    "projectIds",
			Types:    "select",
			Children: projectList,
		},
		{
			Label:    "集群",
			Value:    "clusterName",
			Types:    "select",
			Children: clusterList,
		},
		{
			Label:    "平台",
			Value:    "platformIds",
			Types:    "select",
			Children: platformList,
		},
		{
			Label: "开服时间",
			Value: "dateRange",
			Types: "daterange",
		},
		{
			Label:    "安装状态",
			Value:    "installStatus",
			Types:    "select",
			Children: fl,
		},
		{
			Label:    "清档状态",
			Types:    "select",
			Value:    "initdbStatus",
			Children: fl2,
		},
	}

	return &types.ListOpenPlanResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
