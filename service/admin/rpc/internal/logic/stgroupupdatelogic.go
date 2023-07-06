package logic

import (
	"context"
	"time"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StgroupUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStgroupUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StgroupUpdateLogic {
	return &StgroupUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StgroupUpdateLogic) StgroupUpdate(in *adminclient.StgroupUpdateReq) (*adminclient.StgroupUpdateResp, error) {
	err := l.svcCtx.StgroupModel.Update(l.ctx, &model.SysStgroup{
		Id:             in.Id,
		StRemark:       in.StRemark,
		StName:         in.StName,
		StJson:         in.StJson,
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}
	return &adminclient.StgroupUpdateResp{}, nil
}
