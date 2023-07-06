package logic

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClusterListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClusterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterListLogic {
	return &ClusterListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Platform Rpc End
func (l *ClusterListLogic) ClusterList(in *yunweiclient.ListClusterReq) (*yunweiclient.ListClusterResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "view_label_id__=", in.LabelId,
		"view_project_id__=", in.ProjectId,
		"view_project_id__in", in.ProjectIds,
		"view_label_id__in", in.ClusterCn,
	)

	list, err := l.svcCtx.PlatformModel.FindListByClusterId(l.ctx, in.Current, in.PageSize, filters...)

	if err != nil {
		return nil, xerr.NewErrMsg("查询集群信息失败，原因：" + err.Error())
	}
	count, err := l.svcCtx.PlatformModel.CountCluster(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("统计集群信息失败，原因：" + err.Error())

	}

	var tmp []*yunweiclient.ListClusterData
	err = copier.Copy(&tmp, list)
	if err != nil {
		return nil, xerr.NewErrMsg("复制集群信息失败，原因：" + err.Error())
	}
	tmpPlatform := make([]*yunweiclient.ListPlatformInfoData, 0)
	tmpAsset := make([]*yunweiclient.ListAssetInfoData, 0)
	if in.LabelId != 0 && in.ProjectId != 0 {
		filters = make([]interface{}, 0)
		filters = append(filters, "label_id__=", in.LabelId,
			"project_id__=", in.ProjectId,
			"json_unquote(view_json_id->'$.view_recycle_type')__=", gconv.Int64(2),
		)
		assetList, err := l.svcCtx.PlatformModel.FindAssetListByClusterId(l.ctx, filters...)
		if err != nil {
			return nil, xerr.NewErrMsg("查询集群资产信息失败，原因：" + err.Error())
		}

		err = copier.Copy(&tmpAsset, assetList)
		if err != nil {
			return nil, xerr.NewErrMsg("复制集群资产信息失败，原因：" + err.Error())
		}

		platformList, err := l.svcCtx.PlatformModel.FindPlatformListByClusterId(l.ctx, filters...)
		if err != nil {
			return nil, xerr.NewErrMsg("查询集群平台信息失败，原因：" + err.Error())
		}
		err = copier.Copy(&tmpPlatform, platformList)
		if err != nil {
			return nil, xerr.NewErrMsg("复制集群平台信息失败，原因：" + err.Error())
		}
	}
	return &yunweiclient.ListClusterResp{
		Rows:         tmp,
		Total:        count,
		AssetRows:    tmpAsset,
		PlatformRows: tmpPlatform,
	}, nil
}
