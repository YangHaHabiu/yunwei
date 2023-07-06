package insideTasks

import (
	"context"
	"ywadmin-v3/service/intranet/api/internal/logic/common"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideTasksListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksListLogic {
	return &InsideTasksListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideTasksListLogic) InsideTasksList(req *types.ListInsideTasksReq) (resp *types.ListInsideTasksResp, err error) {
	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListInsideTasksData, 0)
	list, err := l.svcCtx.IntranetRpc.InsideTasksList(l.ctx, &intranetclient.ListInsideTasksReq{
		Current:      req.Current,
		PageSize:     req.PageSize,
		ProjectId:    req.ProjectId,
		ProjectIds:   projectIds,
		VersionId:    req.VersionId,
		ServerId:     req.ServerId,
		TasksType:    req.TasksType,
		RecentSubmit: req.RecentSubmit,
		Status:       req.Status,
		StartTime:    req.StartTime,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	itsList, err := common.GetinsideTasksStatus(l.svcCtx, l.ctx)
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
			Label:    "任务状态",
			Value:    "status",
			Types:    "select",
			Children: itsList,
		},
		{
			Label: "创建时间",
			Value: "startTime",
			Types: "daterange",
		},
	}

	return &types.ListInsideTasksResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
