package company

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompanyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyListLogic {
	return &CompanyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompanyListLogic) CompanyList(req *types.ListCompanyReq) (*types.ListCompanyResp, error) {
	resp, err := l.svcCtx.AdminRpc.CompanyList(l.ctx, &admin.CompanyListReq{
		Current:             req.Current,
		PageSize:            req.PageSize,
		CompanyCn:           req.CompanyCn,
		SupplyCompanyStatus: req.SupplyCompanyStatus,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询机构列表异常:%s", string(data), err.Error())
		return nil, err
	}
	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label: "公司名",
			Value: "companyCn",
			Types: "input",
		},
		{
			Label: "出机方状态",
			Types: "select",
			Value: "supplyCompanyStatus",
			Children: []*types.FilterList{
				{
					Label: "显示",
					Value: "1",
				},
				{
					Label: "隐藏",
					Value: "2",
				},
			},
		},
	}

	list := make([]*types.ListCompanyData, 0)

	for _, data := range resp.List {
		list = append(list, &types.ListCompanyData{
			CompanyId:           data.CompanyId,
			CompanyEn:           data.CompanyEn,
			CompanyCn:           data.CompanyCn,
			SupplyCompanyStatus: data.SupplyCompanyStatus,
		})
	}

	return &types.ListCompanyResp{
		Rows:   list,
		Total:  resp.Total,
		Filter: filterList,
	}, nil

}
