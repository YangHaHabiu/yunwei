package taskQueue

import (
	"context"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebSocketAddrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWebSocketAddrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebSocketAddrLogic {
	return &GetWebSocketAddrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebSocketAddrLogic) GetWebSocketAddr() (resp *types.GetWebSocketAddrResp, err error) {
	resp = new(types.GetWebSocketAddrResp)
	resp.IpAddr = l.svcCtx.Config.WebsocketAddr
	return
}
