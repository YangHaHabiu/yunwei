package configFileDelivery

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileDeliveryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryAddLogic {
	return &ConfigFileDeliveryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileDeliveryAddLogic) ConfigFileDeliveryAdd(req *types.AddConfigFileDeliveryReq) error {
	var tmp []*yunweiclient.AddConfigFileDeliveryDataList
	err := copier.Copy(&tmp, req.ConfigFileData)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.ConfigFileDeliveryAdd(l.ctx, &yunweiclient.AddConfigFileDeliveryReq{ConfigFileData: tmp})
	if err != nil {
		return err
	}
	return nil
}
