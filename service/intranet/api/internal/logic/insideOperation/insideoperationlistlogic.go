package insideOperation

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideOperationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideOperationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideOperationListLogic {
	return &InsideOperationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideOperationListLogic) InsideOperationList(req *types.ListInsideOperationReq) (resp *types.ListInsideOperationResp, err error) {
	tmp := make([]*types.ListInsideOperationData, 0)
	list, err := l.svcCtx.IntranetRpc.InsideOperationList(l.ctx, &intranetclient.ListInsideOperationReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		ProjectId: req.ProjectId,
		OperType:  req.OperType,
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

	return &types.ListInsideOperationResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
