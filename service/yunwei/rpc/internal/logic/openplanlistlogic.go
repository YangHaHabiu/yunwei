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

type OpenPlanListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenPlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanListLogic {
	return &OpenPlanListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenPlanListLogic) OpenPlanList(in *yunweiclient.ListOpenPlanReq) (*yunweiclient.ListOpenPlanResp, error) {
	var (
		count    int64
		list     []*yunweiclient.ListOpenPlanData
		all      *[]model.OpenPlanList
		err      error
		orderTmp string
	)

	filters := make([]interface{}, 0)
	filters = append(filters,
		"open_time__xrange", in.DateRange,
		"project_id__in", in.ProjectIds,
		"platform_ex__in", in.PlatformIds,
		"cluster_id__in", in.ClusterName,
		"install_status__in", in.InstallStatus,
		"initdb_status__in", in.InitdbStatus,
	)

	if in.GameType == "cross" {
		filters = append(filters, "platform_en__rlike", "cross")
	} else if in.GameType == "notCross" {
		filters = append(filters, "platform_en__nrlike", "cross")
	}

	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.OpenPlanModel.FindAll(l.ctx, filters...)
		if err != nil {
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

		count, _ = l.svcCtx.OpenPlanModel.Count(l.ctx, filters...)

		all, err = l.svcCtx.OpenPlanModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, orderTmp, filters...)
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
	return &yunweiclient.ListOpenPlanResp{
		Rows:  list,
		Total: count,
	}, nil
}
