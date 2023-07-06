package configFile

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileAddLogic {
	return &ConfigFileAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileAddLogic) ConfigFileAdd(req *types.AddConfigFileReq) error {
	var tmp yunweiclient.ConfigFileCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.ConfigFileAdd(l.ctx, &yunweiclient.AddConfigFileReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil

}
