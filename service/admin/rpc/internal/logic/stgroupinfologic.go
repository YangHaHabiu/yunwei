package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StgroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStgroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StgroupInfoLogic {
	return &StgroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StgroupInfoLogic) StgroupInfo(in *adminclient.StgroupInfoReq) (*adminclient.StgroupInfoResp, error) {
	info, err := l.svcCtx.StgroupModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, xerr.NewErrCode(xerr.ADMIN_NOTFOUNDUID_ERROR)
	default:
		return nil, err
	}
	return &adminclient.StgroupInfoResp{Info: &adminclient.StgroupListData{
		Id:             info.Id,
		StName:         info.StName,
		StRemark:       info.StRemark,
		StJson:         info.StJson,
		CreateBy:       info.CreateBy,
		CreateTime:     info.CreateTime.Format("2006-01-02 15:04:05"),
		LastUpdateBy:   info.LastUpdateBy,
		LastUpdateTime: info.LastUpdateTime.Format("2006-01-02 15:04:05"),
		DelFlag:        info.DelFlag,
	}}, nil
}
