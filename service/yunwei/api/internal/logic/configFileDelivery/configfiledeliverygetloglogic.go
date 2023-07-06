package configFileDelivery

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
	"ywadmin-v3/service/yunwei/api/internal/svc"
)

type ConfigFileDeliveryGetLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileDeliveryGetLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryGetLogLogic {
	return &ConfigFileDeliveryGetLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileDeliveryGetLogLogic) ConfigFileDeliveryGetLog() (err error) {

	_, err = l.svcCtx.YunWeiRpc.ConfigFileDeliveryGetLog(l.ctx, &yunweiclient.GetConfigFileDeliveryLogReq{})
	if err != nil {
		return err
	}

	return nil
}
