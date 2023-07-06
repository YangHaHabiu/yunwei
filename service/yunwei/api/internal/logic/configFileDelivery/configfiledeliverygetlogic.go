package configFileDelivery

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileDeliveryGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryGetLogic {
	return &ConfigFileDeliveryGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileDeliveryGetLogic) ConfigFileDeliveryGet(req *types.GetConfigFileDeliveryTreeReq) (resp *types.GetConfigFileDeliveryTreeResp, err error) {

	var tmp yunweiclient.GetConfigFileDeliveryTreeReq
	err = copier.Copy(&tmp, req)
	if err != nil {
		return nil, xerr.NewErrMsg("新增拷贝数据失败，原因：" + err.Error())
	}
	content, err := l.svcCtx.YunWeiRpc.ConfigFileDeliveryGet(l.ctx, &tmp)
	if err != nil {
		return nil, err
	}
	resp = new(types.GetConfigFileDeliveryTreeResp)
	resp.Rows = content.Rows

	return
}
