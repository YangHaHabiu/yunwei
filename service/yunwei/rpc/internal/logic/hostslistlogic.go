package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"strings"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type HostsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHostsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HostsListLogic {
	return &HostsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OpenPlan Rpc End
func (l *HostsListLogic) HostsList(in *yunweiclient.ListHostsReq) (*yunweiclient.ListHostsResp, error) {
	filters := make([]interface{}, 0)
	viewHostRoleCn := strings.ReplaceAll(in.ViewHostRoleCn, ",", "|")
	filters = append(filters, "view_host_role_cn__REGEXP", viewHostRoleCn,
		"view_user_project_id__in", in.ProjectIds,
		"view_asset_ownership_company_id__in", in.Company,
		"view_outer_ip@view_inner_ip__or__regexp", in.Ips,
		"view_provider_id__in", in.Provider,
		"label_names__like", in.Label,
	)
	list, err := l.svcCtx.PlatformModel.FindPageHostsListByPage(l.ctx, in.Current, in.PageSize, in.SNames, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询服务器信息失败，原因：" + err.Error())
	}
	count, err := l.svcCtx.PlatformModel.CountHosts(l.ctx, in.SNames, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("统计服务器信息失败，原因：" + err.Error())
	}
	var tmp []*yunweiclient.ListHostsData
	err = copier.Copy(&tmp, list)
	if err != nil {
		return nil, xerr.NewErrMsg("复制服务器信息失败，原因：" + err.Error())
	}
	return &yunweiclient.ListHostsResp{
		Rows:  tmp,
		Total: count,
	}, nil
}
