package taskQueue

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTasksListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksListLogic {
	return &TasksListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TasksListLogic) TasksList(req *types.ListTasksReq) (resp *types.ListTasksResp, err error) {

	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	//用户列表
	userList, err := common.GetUserList(l.svcCtx, l.ctx, projectIds)
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.ListTasksData, 0)
	list, err := l.svcCtx.YunWeiRpc.TasksList(l.ctx, &yunweiclient.ListTasksReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		ProjectIds: projectIds,
		TaskType:   req.TaskType,
		CreateTime: req.CreateTime,
		CreateBy:   req.CreateBy,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
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
			Label: "任务类型",
			Value: "taskType",
			Types: "select",
			Children: []*types.FilterList{
				{
					Label: "临时维护",
					Value: "1",
				},
				{
					Label: "日常维护",
					Value: "2",
				},
			},
		},
		{
			Label:    "用户",
			Value:    "createBy",
			Types:    "select",
			Children: userList,
		},
		{
			Label: "创建时间",
			Value: "createTime",
			Types: "daterange",
		},
	}

	return &types.ListTasksResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
