package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UgroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUgroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupInfoLogic {
	return &UgroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UgroupInfoLogic) UgroupInfo(in *adminclient.UgroupInfoReq) (*adminclient.UgroupInfoResp, error) {
	info, err := l.svcCtx.UgroupModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, xerr.NewErrCode(xerr.ADMIN_NOTFOUNDUID_ERROR)
	default:
		return nil, err
	}
	return &adminclient.UgroupInfoResp{Info: &adminclient.UgroupListData{
		Id:             info.Id,
		UgName:         info.Ugname,
		UgJson:         info.UgJson,
		CreateBy:       info.CreateBy,
		CreateTime:     info.CreateTime.Format("2006-01-02 15:04:05"),
		LastUpdateBy:   info.LastUpdateBy,
		LastUpdateTime: info.LastUpdateTime.Format("2006-01-02 15:04:05"),
		DelFlag:        info.DelFlag,
	}}, nil
}
