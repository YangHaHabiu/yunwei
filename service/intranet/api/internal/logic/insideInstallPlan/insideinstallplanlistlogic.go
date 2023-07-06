package insideInstallPlan

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideInstallPlanListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideInstallPlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideInstallPlanListLogic {
	return &InsideInstallPlanListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideInstallPlanListLogic) InsideInstallPlanList(req *types.ListInsideInstallPlanReq) (resp *types.ListInsideInstallPlanResp, err error) {
	tmp := make([]*types.ListInsideInstallPlanData, 0)
	list, err := l.svcCtx.IntranetRpc.InsideInstallPlanList(l.ctx, &intranetclient.ListInsideInstallPlanReq{
		Current:  req.Current,
		PageSize: req.PageSize,
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

	return &types.ListInsideInstallPlanResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
