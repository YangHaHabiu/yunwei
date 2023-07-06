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

type ConfigFileGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileGetLogic {
	return &ConfigFileGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileGetLogic) ConfigFileGet(req *types.GetConfigFileReq) (resp *types.ListConfigFileData, err error) {
	get, err := l.svcCtx.YunWeiRpc.ConfigFileGet(l.ctx, &yunweiclient.GetConfigFileReq{ConfigFileId: req.ConfigFileId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListConfigFileData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
