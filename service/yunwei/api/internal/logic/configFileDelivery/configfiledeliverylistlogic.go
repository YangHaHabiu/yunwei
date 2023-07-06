package configFileDelivery

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileDeliveryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryListLogic {
	return &ConfigFileDeliveryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileDeliveryListLogic) ConfigFileDeliveryList(req *types.ListConfigFileDeliveryReq) (resp *types.ListConfigFileDeliveryResp, err error) {
	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	list, err := l.svcCtx.YunWeiRpc.ConfigFileDeliveryList(l.ctx, &yunweiclient.ListConfigFileDeliveryReq{
		ProjectIds: projectIds,
	})
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.ListConfigFileData, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}
	tmp1 := make([]*types.ListConfigFileDeliveryDataTree, 0)

	for _, v := range list.MergeRows {
		t := make([]string, 0)
		copier.Copy(&t, v.MouldFile)

		tmp1 = append(tmp1, &types.ListConfigFileDeliveryDataTree{
			ProjectId: v.ProjectId,
			TotalList: v.TotalList,
			MouldFile: t,
		})

	}

	//err = copier.Copy(&tmp1, list.MergeRows)
	//if err != nil {
	//	return nil, err
	//}

	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label:    "项目",
			Value:    "projectIds",
			Types:    "select",
			Children: projectList,
		},
	}

	resp = new(types.ListConfigFileDeliveryResp)
	resp.Rows = tmp
	resp.MergeRows = tmp1
	resp.Filter = filterList

	return
}
