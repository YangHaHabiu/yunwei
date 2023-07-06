package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMaintainPlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanListLogic {
	return &MaintainPlanListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MaintainPlanListLogic) MaintainPlanList(in *yunweiclient.ListMaintainPlanReq) (*yunweiclient.ListMaintainPlanResp, error) {
	var (
		count    int64
		list     []*yunweiclient.ListMaintainPlanData
		all      *[]model.MaintainPlanList
		err      error
		orderTmp string
	)

	filters := make([]interface{}, 0)
	filters = append(filters,
		"project_id__=", in.ProjectId,
		"start_time__>=", in.StartTime,
		"start_time__=", in.PlanStartTime,
		"start_time__xrange", in.DateRange,
		"maintain_type__in", in.MaintainType,
		"project_id__in", in.ProjectIds,
		"title__like", in.Title,
		"task_id__=", in.TaskId,
		"id__=", in.MaintainId,
	)
	if in.Current == 0 && in.PageSize == 0 {
		all, err = l.svcCtx.MaintainPlanModel.FindAll(l.ctx, filters...)
		if err != nil {
			reqStr, _ := json.Marshal(in)
			logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
			return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
		}
	} else {
		sortList := make([]string, 0)
		for _, v := range in.SortFiledList {
			sortList = append(sortList, fmt.Sprintf("%s %s", tool.Camel2Case(v.Field), v.Order))
		}
		if len(sortList) != 0 {
			orderTmp = fmt.Sprintf("%s %s", "order by", strings.Join(sortList, ","))
		}
		count, _ = l.svcCtx.MaintainPlanModel.Count(l.ctx, filters...)
		all, err = l.svcCtx.MaintainPlanModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, orderTmp, filters...)
		if err != nil {
			reqStr, _ := json.Marshal(in)
			logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
			return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
		}
	}

	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	return &yunweiclient.ListMaintainPlanResp{
		Rows:  list,
		Total: count,
	}, nil
}
