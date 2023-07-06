package insideProjectCluster

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProjectClusterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideProjectClusterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProjectClusterListLogic {
	return &InsideProjectClusterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideProjectClusterListLogic) InsideProjectClusterList(req *types.ListInsideProjectClusterReq) (resp *types.ListInsideProjectClusterResp, err error) {
	tmp := make([]*types.ListInsideProjectClusterData, 0)
	list, err := l.svcCtx.IntranetRpc.InsideProjectClusterList(l.ctx, &intranetclient.ListInsideProjectClusterReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		ProjectId: req.ProjectId,
		ClusterId: req.ClusterId,
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

	return &types.ListInsideProjectClusterResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil

}
