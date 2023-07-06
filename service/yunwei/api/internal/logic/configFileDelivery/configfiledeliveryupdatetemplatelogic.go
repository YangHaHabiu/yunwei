package configFileDelivery

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
	"ywadmin-v3/service/yunwei/rpc/yunwei"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryUpdateTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileDeliveryUpdateTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryUpdateTemplateLogic {
	return &ConfigFileDeliveryUpdateTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileDeliveryUpdateTemplateLogic) ConfigFileDeliveryUpdateTemplate(req *types.UpdateConfigFileDeliveryTemplateReq) error {
	var tmp yunwei.UpdateConfigFileDeliveryTemplateReq
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("更新拷贝数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.ConfigFileDeliveryUpdateTemplate(l.ctx, &tmp)
	if err != nil {
		return err
	}
	return nil
}
