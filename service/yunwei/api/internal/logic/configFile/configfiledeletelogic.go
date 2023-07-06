package configFile

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeleteLogic {
	return &ConfigFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigFileDeleteLogic) ConfigFileDelete(req *types.DeleteConfigFileReq) error {
	_, err := l.svcCtx.YunWeiRpc.ConfigFileDelete(l.ctx, &yunweiclient.DeleteConfigFileReq{ConfigFileId: req.ConfigFileId})
	if err != nil {
		return err
	}
	return nil
}
