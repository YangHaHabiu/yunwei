package company

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompanyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyUpdateLogic {
	return &CompanyUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompanyUpdateLogic) CompanyUpdate(req *types.UpdateCompanyReq) (err error) {
	_, err = l.svcCtx.AdminRpc.CompanyUpdate(l.ctx, &admin.CompanyUpdateReq{
		CompanyId: req.CompanyId,
		CompanyCn: req.CompanyCn,
		CompanyEn: req.CompanyEn,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return
}
