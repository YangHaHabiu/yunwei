package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyListLogic {
	return &CompanyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyListLogic) CompanyList(in *adminclient.CompanyListReq) (*adminclient.CompanyListResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters,
		"company_cn@company_en__or__regexp", in.CompanyCn,
		"supply_company_status__in", in.SupplyCompanyStatus,
	)

	var (
		all   *[]model.Company
		err   error
		count int64
	)

	if in.PageSize == 0 && in.Current == 0 {
		all, err = l.svcCtx.CompanyModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.CompanyModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
		count, _ = l.svcCtx.CompanyModel.Count(l.ctx, filters...)
	}

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrCode(xerr.ADMIN_DEPTSELECT_ERROR)
	}

	var list []*adminclient.CompanyListData
	for _, data := range *all {
		list = append(list, &adminclient.CompanyListData{
			CompanyEn:           data.CompanyEn,
			CompanyCn:           data.CompanyCn,
			CompanyId:           data.CompanyId,
			SupplyCompanyStatus: data.SupplyCompanyStatus,
		})
	}

	return &adminclient.CompanyListResp{
		List:  list,
		Total: count,
	}, nil
}
