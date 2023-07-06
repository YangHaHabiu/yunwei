package maintainPlan

import (
	"context"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaintainPlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanListLogic {
	return &MaintainPlanListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaintainPlanListLogic) MaintainPlanList(req *types.ListMaintainPlanReq) (resp *types.ListMaintainPlanResp, err error) {

	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
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
			Label: "标题",
			Value: "title",
			Types: "input",
		},
		{
			Label: "开始时间",
			Value: "dateRange",
			Types: "daterange",
		},
	}

	sortTmpList := make([]*yunweiclient.SortFiledData, 0)
	err = copier.Copy(&sortTmpList, req.SortFields)
	if err != nil {
		return nil, err
	}

	//查询维护计划
	tmp := make([]*types.ListMaintainPlanData, 0)
	list, err := l.svcCtx.YunWeiRpc.MaintainPlanList(l.ctx, &yunweiclient.ListMaintainPlanReq{
		Current:       req.Current,
		PageSize:      req.PageSize,
		ProjectIds:    projectIds,
		DateRange:     req.DateRange,
		Title:         req.Title,
		SortFiledList: sortTmpList,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}
	return &types.ListMaintainPlanResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil

}
