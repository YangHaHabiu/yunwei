package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeatureServerInfoGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeatureServerInfoGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeatureServerInfoGetLogic {
	return &FeatureServerInfoGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeatureServerInfoGetLogic) FeatureServerInfoGet(in *yunweiclient.GetFeatureServerReq) (*yunweiclient.ListFeatureServerData, error) {
	one, err := l.svcCtx.FeatureServerInfoModel.FindOne(l.ctx, in.FeatureServerId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp yunweiclient.ListFeatureServerData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
