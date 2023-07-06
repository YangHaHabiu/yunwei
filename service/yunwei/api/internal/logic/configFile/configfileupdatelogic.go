package configFile

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

type ConfigFileUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileUpdateLogic {
	return &ConfigFileUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileUpdateLogic) ConfigFileUpdate(req *types.UpdateConfigFileReq) error {
	var tmp yunwei.ConfigFileCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("更新拷贝数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.ConfigFileUpdate(l.ctx, &yunweiclient.UpdateConfigFileReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
