package keyManage

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KeyManageDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeyManageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeyManageDeleteLogic {
	return &KeyManageDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KeyManageDeleteLogic) KeyManageDelete(req *types.DeleteKeyManageReq) error {
	_, err := l.svcCtx.YunWeiRpc.KeyManageDelete(l.ctx, &yunweiclient.DeleteKeyManageReq{KeyId: req.KeyId})
	if err != nil {
		return err
	}
	return nil
}
