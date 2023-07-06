package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyDeleteLogic {
	return &CompanyDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyDeleteLogic) CompanyDelete(in *adminclient.CompanyDeleteReq) (*adminclient.CompanyDeleteResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "company_id__=", in.CompanyId)
	//判断关联的项目
	all, err := l.svcCtx.ProjectRelationshipModel.FindAll(l.ctx, filters...)
	if err != nil || len(*all) != 0 {
		return nil, xerr.NewErrMsg("公司关联项目数据，禁止删除，请检查")
	}

	//判断关联的机器
	allx, err := l.svcCtx.ServerAffiliationModel.FindAll(l.ctx, filters...)
	if err != nil || len(*allx) != 0 {
		return nil, xerr.NewErrMsg("公司关联机器数据，禁止删除，请检查")
	}

	err = l.svcCtx.CompanyModel.DeleteSoft(l.ctx, in.CompanyId)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_DELETE_ERROR)
	}
	return &adminclient.CompanyDeleteResp{}, nil
}
