package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"strings"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type WaitInstallListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWaitInstallListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WaitInstallListLogic {
	return &WaitInstallListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WaitInstallListLogic) WaitInstallList(in *yunweiclient.ListWaitInstallReq) (*yunweiclient.ListWaitInstallResp, error) {
	var (
		orderTmp string
	)
	filters := make([]interface{}, 0)
	filters = append(filters,
		"install_status__=", int64(-1),
		"project_id__in", in.ProjectIds,
	)

	sortList := make([]string, 0)
	for _, v := range in.SortFiledList {
		sortList = append(sortList, fmt.Sprintf("%s %s", tool.Camel2Case(v.Field), v.Order))
	}
	if len(sortList) != 0 {
		orderTmp = fmt.Sprintf("%s %s", "order by", strings.Join(sortList, ","))
	}
	count, err := l.svcCtx.OpenPlanModel.Count(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("统计待装服失败，原因：" + err.Error())
	}

	page, err := l.svcCtx.OpenPlanModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, orderTmp, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询待装服列表失败，原因：" + err.Error())
	}
	var tmp []*yunweiclient.ListWaitInstallData
	err = copier.Copy(&tmp, page)
	if err != nil {
		return nil, xerr.NewErrMsg("复制待装服列表失败，原因：" + err.Error())
	}

	return &yunweiclient.ListWaitInstallResp{
		Rows:  tmp,
		Total: count,
	}, nil
}
