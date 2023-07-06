package insideServer

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideServerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideServerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideServerListLogic {
	return &InsideServerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideServerListLogic) InsideServerList(req *types.ListInsideServerReq) (resp *types.ListInsideServerResp, err error) {
	tmp := make([]*types.ListInsideServerData, 0)
	list, err := l.svcCtx.IntranetRpc.InsideServerList(l.ctx, &intranetclient.ListInsideServerReq{
		Current:     req.Current,
		PageSize:    req.PageSize,
		FeatureType: req.FeatureType,
		ProjectId:   req.ProjectId,
		ClusterId:   req.ClusterId,
		BuildType:   req.BuildType,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label: "test1",
			Value: "test1",
			Types: "input",
		},
		{
			Label: "time",
			Value: "test2",
			Types: "daterange",
		},
		{
			Label: "test3",
			Value: "test",
			Types: "select",
			Children: []*types.FilterList{
				{
					Label: "l1",
					Value: "1",
				},
			},
		},
	}

	return &types.ListInsideServerResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
