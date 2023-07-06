package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddLogic {
	return &CompanyAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// company rpc start
func (l *CompanyAddLogic) CompanyAdd(in *adminclient.CompanyAddReq) (*adminclient.CompanyAddResp, error) {
	_, err := l.svcCtx.CompanyModel.Insert(l.ctx, &model.Company{
		CompanyCn: in.CompanyCn,
		CompanyEn: in.CompanyEn,
		DelFlag:   0,
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
	}

	return &adminclient.CompanyAddResp{}, nil
}
