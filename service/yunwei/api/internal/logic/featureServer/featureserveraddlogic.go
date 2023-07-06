package featureServer

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunwei"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeatureServerAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeatureServerAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeatureServerAddLogic {
	return &FeatureServerAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeatureServerAddLogic) FeatureServerAdd(req *types.AddFeatureServerReq) error {
	var tmp yunwei.FeatureServerDatas
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("拷贝数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.FeatureServerAdd(l.ctx, &yunweiclient.AddFeatureServerReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
