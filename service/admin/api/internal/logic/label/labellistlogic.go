package label

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/api/internal/logic/common"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LabelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLabelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LabelListLogic {
	return &LabelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LabelListLogic) LabelList(req *types.ListLabelReq) (*types.ListLabelResp, error) {
	resp, err := l.svcCtx.AdminRpc.LabelList(l.ctx, &admin.LabelListReq{
		Current:     req.Current,
		PageSize:    req.PageSize,
		LabelName:   req.LabelName,
		LabelId:     req.LabelId,
		LabelType:   req.LabelType,
		LabelValues: req.LabelValues,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询列表异常:%s", string(data), err.Error())
		return nil, err
	}

	list := make([]*types.ListLabelData, 0)

	for _, data := range resp.List {
		list = append(list, &types.ListLabelData{
			ViewLabelId:     data.ViewLabelId,
			ViewLabelValues: data.ViewLabelValues,
			ViewLabelName:   data.ViewLabelName,
			ViewLabelRemark: data.ViewLabelRemark,
			ViewLabelType:   data.ViewLabelType,
			ViewStopStatus:  data.ViewStopStatus,
		})
	}

	byTypes, err := common.GetDictListByTypes(l.svcCtx, l.ctx, "label_type", "标签类型")
	if err != nil {
		return nil, err
	}
	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label: "标签名",
			Value: "labelName",
			Types: "input",
		},
		{
			Label: "标签值",
			Value: "labelValues",
			Types: "input",
		},
		{
			Label:    "标签类型",
			Value:    "labelType",
			Types:    "select",
			Children: byTypes,
		},
	}

	return &types.ListLabelResp{
		Rows:   list,
		Total:  resp.Total,
		Filter: filterList,
	}, nil

}
