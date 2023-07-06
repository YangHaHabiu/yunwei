package platform

import (
	"context"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlatformListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlatformListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformListLogic {
	return &PlatformListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlatformListLogic) PlatformList(req *types.ListPlatformReq) (resp *types.ListPlatformResp, err error) {

	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.ListPlatformData, 0)

	list, err := l.svcCtx.YunWeiRpc.PlatformList(l.ctx, &yunweiclient.ListPlatformReq{
		Current:      req.Current,
		PageSize:     req.PageSize,
		ProjectIds:   projectIds,
		Id:           req.Id,
		PlatformInfo: req.PlatformInfo,
		Label:        req.Label,
		PlatformType: req.PlatformType,
	})

	if err != nil {
		return nil, err
	}

	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}
	//获取平台
	platformList, err := common.GetPlatform(l.svcCtx, l.ctx, projectIds, "", req.PlatformType)
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
			Label:    "平台",
			Value:    "platformInfo",
			Types:    "select",
			Children: platformList,
		},
		{
			Label: "标签",
			Value: "label",
			Types: "input",
		},
	}

	return &types.ListPlatformResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
