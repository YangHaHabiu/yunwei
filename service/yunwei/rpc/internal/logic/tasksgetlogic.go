package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksGetLogic {
	return &TasksGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TasksGetLogic) TasksGet(in *yunweiclient.GetTasksReq) (*yunweiclient.GetTasksResp, error) {
	one, err := l.svcCtx.TasksModel.FindAll(l.ctx, "id__=", in.Id)
	if err != nil || len(*one) != 1 {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	task := (*one)[0]
	taskList := make([]*model.TasksList, 0)
	taskList = append(taskList, &task)
	allIdList := make([]int64, 0)
	errIdList := make([]int64, 0)

	allIdList = append(allIdList, task.Id)

	oneList, err := l.svcCtx.TasksModel.FindAll(l.ctx, "pid__=", task.Id, "task_status__in", "-1,1,2,3")
	if err != nil {
		return nil, xerr.NewErrMsg("查询一级任务信息失败，原因:" + err.Error())
	}
	for i := 0; i <= len(*oneList)-1; i++ {
		taskList = append(taskList, &(*oneList)[i])
		allIdList = append(allIdList, (*oneList)[i].Id)
		twoList, err := l.svcCtx.TasksModel.FindAll(l.ctx, "pid__=", (*oneList)[i].Id, "task_status__in", "-1,1,2,3")
		if err != nil {
			return nil, xerr.NewErrMsg("查询一级任务信息失败，原因:" + err.Error())
		}
		if (*oneList)[i].TaskStatus == 2 {
			errIdList = append(errIdList, (*oneList)[i].Id)
		}
		for i := 0; i <= len(*twoList)-1; i++ {
			taskList = append(taskList, &(*twoList)[i])
			allIdList = append(allIdList, (*twoList)[i].Id)
			if (*twoList)[i].TaskStatus == 2 {
				errIdList = append(errIdList, (*twoList)[i].Id)
			}
		}
	}

	var tmp []*yunweiclient.ListTasksData
	err = copier.Copy(&tmp, taskList)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}
	return &yunweiclient.GetTasksResp{
		Rows:     tmp,
		AllIdArr: allIdList,
		ErrArr:   errIdList,
	}, nil
}
