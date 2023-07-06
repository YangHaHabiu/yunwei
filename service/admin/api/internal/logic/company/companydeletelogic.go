package company

import (
	"context"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompanyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyDeleteLogic {
	return &CompanyDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompanyDeleteLogic) CompanyDelete(req *types.DeleteCompanyReq) (err error) {
	_, err = l.svcCtx.AdminRpc.CompanyDelete(l.ctx, &admin.CompanyDeleteReq{
		CompanyId: req.CompanyId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除异常:%s", req.CompanyId, err.Error())
		return err
	}
	return
}
