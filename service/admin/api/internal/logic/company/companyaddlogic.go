package company

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompanyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddLogic {
	return &CompanyAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompanyAddLogic) CompanyAdd(req *types.AddCompanyReq) (err error) {
	_, err = l.svcCtx.AdminRpc.CompanyAdd(l.ctx, &admin.CompanyAddReq{
		CompanyCn: req.CompanyCn,
		CompanyEn: req.CompanyEn,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return
}
