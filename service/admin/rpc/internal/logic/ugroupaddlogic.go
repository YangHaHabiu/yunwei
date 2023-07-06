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

type UgroupAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUgroupAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupAddLogic {
	return &UgroupAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ugroup rpc start
func (l *UgroupAddLogic) UgroupAdd(in *adminclient.UgroupAddReq) (*adminclient.UgroupAddResp, error) {
	_, err := l.svcCtx.UgroupModel.Insert(l.ctx, &model.SysUgroup{
		Id:             0,
		UgJson:         in.UgJson,
		Ugname:         in.UgName,
		CreateBy:       in.CreateBy,
		CreateTime:     time.Time{},
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_ADD_ERROR)
	}
	return &adminclient.UgroupAddResp{}, nil
}
