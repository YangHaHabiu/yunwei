package hosts

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHostsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostsListLogic {
	return &HostsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HostsListLogic) HostsList(req *types.ListHostsReq) (resp *types.ListHostsResp, err error) {
	//出机方
	companyList, err := common.GetCompanyList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}
	//云商
	providerList, err := common.GetProviderList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err
	}

	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	list, err := l.svcCtx.YunWeiRpc.HostsList(l.ctx, &yunweiclient.ListHostsReq{
		Current:        req.Current,
		PageSize:       req.PageSize,
		ViewHostRoleCn: req.ViewHostRoleCn,
		Company:        req.Company,
		ProjectIds:     projectIds,
		Ips:            req.Ips,
		SNames:         req.SNames,
		Provider:       req.Provider,
		Label:          req.Label,
	})
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListHostsData, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制服务器信息出错，原因：" + err.Error())
	}
	//获取集群
	//clusterList, err := common.GetCluster(l.svcCtx, l.ctx)
	//if err != nil {
	//	return nil, err
	//}
	//自定义筛选条件
	filterList := []*types.FilterList{

		{
			Label:    "出机方",
			Value:    "company",
			Types:    "select",
			Children: companyList,
		},
		{
			Label:    "项目",
			Value:    "projectIds",
			Types:    "select",
			Children: projectList,
		},
		{
			Label: "IP",
			Value: "ips",
			Types: "input",
		},
		{
			Label:    "云商",
			Value:    "provider",
			Types:    "select",
			Children: providerList,
		},
		{
			Label: "服名称",
			Value: "sNames",
			Types: "input",
		},

		{
			Label: "标签",
			Value: "label",
			Types: "input",
			//Children: clusterList,
		},
	}
	resp = new(types.ListHostsResp)
	resp.Rows = tmp
	resp.Total = list.Total
	resp.Filter = filterList
	return
}
