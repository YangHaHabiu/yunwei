package tasksScheduleQueue

import (
	"context"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksScheduleQueueListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksScheduleQueueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueListLogic {
	return &TasksScheduleQueueListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksScheduleQueueListLogic) TasksScheduleQueueList(req *types.ListTasksScheduleQueueReq) (resp *types.ListTasksScheduleQueueResp, err error) {
	tmp := make([]*types.ListTasksScheduleQueueData, 0)
	list, err := l.svcCtx.YunWeiRpc.TasksScheduleQueueList(l.ctx, &yunweiclient.ListTasksScheduleQueueReq{
		Current:        req.Current,
		PageSize:       req.PageSize,
		ScheduleStatus: req.ScheduleStatus,
		ScheduleType:   req.ScheduleType,
		ScheduleTitle:  req.ScheduleTitle,
		DateRange:      req.DateRange,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	stl, err := common.GetScheduleTypeList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}
	ssl, err := common.GetScheduleStatusList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err

	}
	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label: "计划标题",
			Value: "scheduleTitle",
			Types: "input",
		},
		{
			Label:    "计划类型",
			Value:    "scheduleType",
			Types:    "select",
			Children: stl,
		},
		{
			Label:    "计划状态",
			Value:    "scheduleStatus",
			Types:    "select",
			Children: ssl,
		},
		{
			Label: "开始时间",
			Value: "dateRange",
			Types: "daterange",
		},
	}

	return &types.ListTasksScheduleQueueResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
