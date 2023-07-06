package mergePlan

import (
	"context"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergePlanListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergePlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanListLogic {
	return &MergePlanListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergePlanListLogic) MergePlanList(req *types.ListMergePlanReq) (resp *types.ListMergePlanResp, err error) {
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
	tmp := make([]*types.ListMergePlanData, 0)
	list, err := l.svcCtx.YunWeiRpc.MergePlanList(l.ctx, &yunweiclient.ListMergePlanReq{
		Current:       req.Current,
		PageSize:      req.PageSize,
		ProjectIds:    projectIds,
		MergeStatus:   req.MergeStatus,
		DateRange:     req.DateRange,
		PlatformIds:   req.PlatformIds,
		SortFiledList: sortTmpList,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}
	//获取平台
	platformList, err := common.GetPlatform(l.svcCtx, l.ctx, projectIds, "cross", req.PlatformType)
	if err != nil {
		return nil, err
	}

	ms, err := common.GetMergeStatusList(l.svcCtx, l.ctx)
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
			Label:    "平台",
			Value:    "platformIds",
			Types:    "select",
			Children: platformList,
		},
		{
			Label: "合服开始时间",
			Value: "dateRange",
			Types: "daterange",
		},
		{
			Label:    "合服状态",
			Value:    "mergeStatus",
			Types:    "select",
			Children: ms,
		},
	}

	return &types.ListMergePlanResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
