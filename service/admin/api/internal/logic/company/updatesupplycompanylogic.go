package company

import (
	"context"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSupplyCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSupplyCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSupplyCompanyLogic {
	return &UpdateSupplyCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSupplyCompanyLogic) UpdateSupplyCompany(req *types.UpdateSupplyCompanyReq) error {
	_, err := l.svcCtx.AdminRpc.UpdateSupplyCompany(l.ctx,
		&adminclient.UpdateSupplyCompanyReq{CompanyId: req.CompanyId})
	if err != nil {
		return err
	}
	return nil
}
