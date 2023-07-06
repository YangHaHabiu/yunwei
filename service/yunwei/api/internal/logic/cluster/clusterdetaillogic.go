package cluster

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClusterDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterDetailLogic {
	return &ClusterDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterDetailLogic) ClusterDetail(req *types.DetailClusterReq) (resp *types.DetailClusterResp, err error) {

	list, err := l.svcCtx.YunWeiRpc.ClusterList(l.ctx, &yunweiclient.ListClusterReq{
		PageSize:  1,
		Current:   1,
		LabelId:   req.LabelId,
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, err
	}
	//fmt.Println(list.Rows)
	if len(list.Rows) != 1 {
		return nil, xerr.NewErrMsg("查询集群信息失败")
	}

	tmpAsset := make([]*types.ListClusterAssetData, 0)
	tmpPlatform := make([]*types.ListClusterPlatformData, 0)

	err = copier.Copy(&tmpAsset, list.AssetRows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制集群资产信息出错，原因：" + err.Error())
	}

	err = copier.Copy(&tmpPlatform, list.PlatformRows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制集群平台信息出错，原因：" + err.Error())
	}

	resp = new(types.DetailClusterResp)
	resp.PlatformRows = tmpPlatform
	resp.AssetRows = tmpAsset
	resp.ViewLabelName = list.Rows[0].ViewLabelName
	resp.ViewClusterFeatureInfo = list.Rows[0].ViewClusterFeatureInfo
	resp.ViewProjectCn = list.Rows[0].ViewProjectCn
	return
}
