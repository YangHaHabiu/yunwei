package featureServer

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeatureServerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeatureServerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeatureServerListLogic {
	return &FeatureServerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeatureServerListLogic) FeatureServerList(req *types.ListFeatureServerReq) (*types.ListFeatureServerResp, error) {

	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.ListFeatureServerData, 0)
	list, err := l.svcCtx.YunWeiRpc.FeatureServerList(l.ctx, &yunweiclient.ListFeatureServerReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		FeatureServerId: req.FeatureServerId,
		Ip:              req.Ip,
		Domain:          req.Domain,
		Feature:         req.Feature,
		ProjectIds:      projectIds,
		Remark:          req.Remark,
	})

	if err != nil {
		return nil, err
	}

	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	serverList, err := common.GetFeatureServerList(l.svcCtx, l.ctx)
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
			Label:    "类型",
			Value:    "feature",
			Types:    "select",
			Children: serverList,
		},
		{
			Label: "域名",
			Value: "domain",
			Types: "input",
		},
		{
			Label: "IP",
			Value: "ip",
			Types: "input",
		},
		{
			Label: "备注",
			Value: "remark",
			Types: "input",
		},
	}
	return &types.ListFeatureServerResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
