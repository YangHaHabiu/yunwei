package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UgroupUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUgroupUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupUpdateLogic {
	return &UgroupUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UgroupUpdateLogic) UgroupUpdate(in *adminclient.UgroupUpdateReq) (*adminclient.UgroupUpdateResp, error) {
	err := l.svcCtx.UgroupModel.Update(l.ctx, &model.SysUgroup{
		Id:           in.Id,
		UgJson:       in.UgJson,
		Ugname:       in.UgName,
		LastUpdateBy: in.LastUpdateBy,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}
	return &adminclient.UgroupUpdateResp{}, nil
}
