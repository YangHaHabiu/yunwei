package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StgroupDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStgroupDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StgroupDeleteLogic {
	return &StgroupDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StgroupDeleteLogic) StgroupDelete(in *adminclient.StgroupDeleteReq) (*adminclient.StgroupDeleteResp, error) {

	filters := make([]interface{}, 0)
	filters = append(filters, "stgroup_id__=", in.Id)
	all, err := l.svcCtx.StgroupUserModel.FindAll(l.ctx, filters...)
	if err != nil || len(*all) != 0 {
		return nil, xerr.NewErrMsg("策略组关联用户数据，禁止删除，请检查")
	}
	findAll, err := l.svcCtx.StgroupUgroupModel.FindAll(l.ctx, filters...)
	if err != nil || len(*findAll) != 0 {
		return nil, xerr.NewErrMsg("策略组关联用户组数据，禁止删除，请检查")
	}

	err = l.svcCtx.StgroupModel.DeleteSoft(l.ctx, &model.SysStgroup{Id: in.Id})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_DELETE_ERROR)
	}
	return &adminclient.StgroupDeleteResp{}, nil
}
