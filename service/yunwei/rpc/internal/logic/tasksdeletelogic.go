package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TasksDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTasksDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TasksDeleteLogic {
	return &TasksDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TasksDeleteLogic) TasksDelete(in *yunweiclient.DeleteTasksReq) (*yunweiclient.TasksCommonResp, error) {
	var status int64
	one, err := l.svcCtx.TasksModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	if one.TaskStatus == 1 || one.TaskStatus == 3 || one.TaskStatus == 4 {
		return nil, xerr.NewErrMsg("执行中和执行成功及已取消的任务不允许删除")
	}

	err = TaskCommon(l.svcCtx, l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	//日常维护需要修改维护计划
	if one.TaskType == "2" {
		err = l.svcCtx.MaintainPlanModel.UpdateByField(l.ctx, []string{
			"task_id",
			"id",
		}, 0, one.MaintainId)
		if err != nil {
			return nil, xerr.NewErrMsg("修改维护计划任务id失败" + err.Error())
		}
	}
	if one.TaskStatus == -1 {
		status = 7
	} else {
		status = 4
	}

	err = l.svcCtx.TasksModel.DeleteSoft(l.ctx, in.Id, status)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &yunweiclient.TasksCommonResp{}, nil
}
