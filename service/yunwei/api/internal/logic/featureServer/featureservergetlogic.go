package featureServer

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeatureServerGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeatureServerGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeatureServerGetLogic {
	return &FeatureServerGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeatureServerGetLogic) FeatureServerGet(req *types.GetFeatureServerReq) (resp *types.ListFeatureServerData, err error) {
	get, err := l.svcCtx.YunWeiRpc.FeatureServerInfoGet(l.ctx, &yunweiclient.GetFeatureServerReq{FeatureServerId: req.FeatureServerId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListFeatureServerData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil

}
