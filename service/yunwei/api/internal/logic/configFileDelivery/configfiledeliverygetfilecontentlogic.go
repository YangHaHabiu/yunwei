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

type ConfigFileDeliveryGetFileContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileDeliveryGetFileContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryGetFileContentLogic {
	return &ConfigFileDeliveryGetFileContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileDeliveryGetFileContentLogic) ConfigFileDeliveryGetFileContent(req *types.ConfigFileDeliveryGetFileContentReq) (resp *types.ConfigFileDeliveryGetFileContentResp, err error) {

	var tmp yunweiclient.ConfigFileDeliveryGetFileContentReq
	err = copier.Copy(&tmp, req)
	if err != nil {
		return nil, xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}
	content, err := l.svcCtx.YunWeiRpc.ConfigFileDeliveryGetFileContent(l.ctx, &tmp)
	if err != nil {
		return nil, err
	}
	resp = new(types.ConfigFileDeliveryGetFileContentResp)
	resp.Content = content.Content

	return
}
