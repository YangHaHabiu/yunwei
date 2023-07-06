package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSupplyCompanyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSupplyCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSupplyCompanyLogic {
	return &UpdateSupplyCompanyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSupplyCompanyLogic) UpdateSupplyCompany(in *adminclient.UpdateSupplyCompanyReq) (*adminclient.UpdateSupplyCompanyResp, error) {
	err := l.svcCtx.CompanyModel.UpdateSupplyCompany(l.ctx, in.CompanyId)
	if err != nil {
		return nil, xerr.NewErrMsg("修改出机方的状态失败，原因：" + err.Error())
	}
	return &adminclient.UpdateSupplyCompanyResp{}, nil
}
