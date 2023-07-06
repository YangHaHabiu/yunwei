package cluster

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

type ClusterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterListLogic {
	return &ClusterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterListLogic) ClusterList(req *types.ListClusterReq) (resp *types.ListClusterResp, err error) {

	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}

	list, err := l.svcCtx.YunWeiRpc.ClusterList(l.ctx, &yunweiclient.ListClusterReq{
		PageSize:   req.PageSize,
		Current:    req.Current,
		LabelId:    req.LabelId,
		ProjectId:  req.ProjectId,
		ClusterCn:  req.ClusterCn,
		ProjectIds: projectIds,
	})
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.ListClusterData, 0)

	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制集群资产信息出错，原因：" + err.Error())
	}
	//获取集群
	clusterList, err := common.GetCluster(l.svcCtx, l.ctx, projectIds)
	if err != nil {
		return nil, err
	}
	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label:    "项目",
			Value:    "projectIds",
			Types:    "select",
			Children: projectList,
		},
		{
			Label:    "集群",
			Value:    "clusterCn",
			Types:    "select",
			Children: clusterList,
		},
	}
	resp = new(types.ListClusterResp)
	resp.Rows = tmp
	resp.Total = list.Total
	resp.Filter = filterList

	return
}
