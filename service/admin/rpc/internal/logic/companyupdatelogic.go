package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyUpdateLogic {
	return &CompanyUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyUpdateLogic) CompanyUpdate(in *adminclient.CompanyUpdateReq) (*adminclient.CompanyUpdateResp, error) {
	err := l.svcCtx.CompanyModel.Update(l.ctx, &model.Company{
		CompanyId: in.CompanyId,
		CompanyEn: in.CompanyEn,
		CompanyCn: in.CompanyCn,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}
	return &adminclient.CompanyUpdateResp{}, nil
}
