package logic

import (
	"context"
	"fmt"
	"strings"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type WaitMergeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWaitMergeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WaitMergeListLogic {
	return &WaitMergeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WaitMergeListLogic) WaitMergeList(in *yunweiclient.ListWaitMergeReq) (*yunweiclient.ListWaitMergeResp, error) {
	var (
		orderTmp string
	)
	filters := make([]interface{}, 0)
	filters = append(filters,
		"merge_status__=", int64(-1),
		"project_id__in", in.ProjectIds,
	)

	sortList := make([]string, 0)
	for _, v := range in.SortFiledList {
		sortList = append(sortList, fmt.Sprintf("%s %s", tool.Camel2Case(v.Field), v.Order))
	}
	if len(sortList) != 0 {
		orderTmp = fmt.Sprintf("%s %s", "order by", strings.Join(sortList, ","))
	}
	count, err := l.svcCtx.MergePlanModel.Count(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("统计待合服失败，原因：" + err.Error())
	}

	page, err := l.svcCtx.MergePlanModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, orderTmp, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询待合服列表失败，原因：" + err.Error())
	}
	var tmp []*yunweiclient.ListWaitMergeData
	err = copier.Copy(&tmp, page)
	if err != nil {
		return nil, xerr.NewErrMsg("复制待合服列表失败，原因：" + err.Error())
	}

	return &yunweiclient.ListWaitMergeResp{
		Rows:  tmp,
		Total: count,
	}, nil
}
