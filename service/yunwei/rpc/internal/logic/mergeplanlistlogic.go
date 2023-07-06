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

type MergePlanListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMergePlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanListLogic {
	return &MergePlanListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MergePlanListLogic) MergePlanList(in *yunweiclient.ListMergePlanReq) (*yunweiclient.ListMergePlanResp, error) {
	var (
		count    int64
		list     []*yunweiclient.ListMergePlanData
		all      *[]model.MergePlanList
		err      error
		orderTmp string
	)

	filters := make([]interface{}, 0)

	filters = append(filters,
		"start_time__xrange", in.DateRange,
		"project_id__in", in.ProjectIds,
		"platform_ex__in", in.PlatformIds,
		"merge_status__in", in.MergeStatus,
	)
	count, _ = l.svcCtx.MergePlanModel.Count(l.ctx, filters...)

	sortList := make([]string, 0)
	for _, v := range in.SortFiledList {
		sortList = append(sortList, fmt.Sprintf("%s %s", tool.Camel2Case(v.Field), v.Order))
	}
	if len(sortList) != 0 {
		orderTmp = fmt.Sprintf("%s %s", "order by ", strings.Join(sortList, ","))
	}

	all, err = l.svcCtx.MergePlanModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, orderTmp, filters...)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}

	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	return &yunweiclient.ListMergePlanResp{
		Rows:  list,
		Total: count,
	}, nil
}
