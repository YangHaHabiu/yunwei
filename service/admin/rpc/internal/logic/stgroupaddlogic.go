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

type StgroupAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStgroupAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StgroupAddLogic {
	return &StgroupAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// stgroup rpc start
func (l *StgroupAddLogic) StgroupAdd(in *adminclient.StgroupAddReq) (*adminclient.StgroupAddResp, error) {
	_, err := l.svcCtx.StgroupModel.Insert(l.ctx, &model.SysStgroup{
		Id:             0,
		StName:         in.StName,
		StJson:         in.StJson,
		StRemark:       in.StRemark,
		CreateBy:       in.CreateBy,
		CreateTime:     time.Now(),
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
	}
	return &adminclient.StgroupAddResp{}, nil
}
