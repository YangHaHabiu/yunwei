package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileDeliveryGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryGetLogic {
	return &ConfigFileDeliveryGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigFileDeliveryGetLogic) ConfigFileDeliveryGet(in *yunweiclient.GetConfigFileDeliveryTreeReq) (*yunweiclient.GetConfigFileDeliveryTreeResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "view_user_project_id__=", in.ProjectId, "view_config_file_id__=", in.ConfigFileId)
	all, err := l.svcCtx.ConfigMngLogModel.FindPageList(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询配置下发服务器信息失败" + err.Error())
	}

	if len(*all) != 1 {
		return nil, xerr.NewErrMsg("查询数据错误")
	}
	return &yunweiclient.GetConfigFileDeliveryTreeResp{
		Rows: (*all)[0].AssetIps,
	}, nil
}
