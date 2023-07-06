package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksScheduleQueueListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksScheduleQueueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksScheduleQueueListLogic {
	return &TasksScheduleQueueListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TasksScheduleQueueListLogic) TasksScheduleQueueList(in *yunweiclient.ListTasksScheduleQueueReq) (*yunweiclient.ListTasksScheduleQueueResp, error) {
	var (
		count int64
		list  []*yunweiclient.ListTasksScheduleQueueData
		all   *[]model.TasksScheduleQueueNew
		err   error
	)

	filters := make([]interface{}, 0)
	filters = append(filters,
		"schedule_type__in", in.ScheduleType,
		"schedule_status__in", in.ScheduleStatus,
		"schedule_start_time__xrange", in.DateRange,
		"schedule_title__like", in.ScheduleTitle,
	)
	count, _ = l.svcCtx.TasksScheduleQueueModel.Count(l.ctx, filters...)
	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.TasksScheduleQueueModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.TasksScheduleQueueModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	}
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}

	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	return &yunweiclient.ListTasksScheduleQueueResp{
		Rows:  list,
		Total: count,
	}, nil
}
