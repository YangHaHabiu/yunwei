package switchEntranceGameserver

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchEntranceGameserverDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwitchEntranceGameserverDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchEntranceGameserverDeleteLogic {
	return &SwitchEntranceGameserverDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwitchEntranceGameserverDeleteLogic) SwitchEntranceGameserverDelete(req *types.DeleteSwitchEntranceGameserverReq) error {
	_, err := l.svcCtx.YunWeiRpc.SwitchEntranceGameserverDelete(l.ctx, &yunweiclient.DeleteSwitchEntranceGameserverReq{Ids: req.Ids, Operation: req.Operation})
	if err != nil {
		return err
	}
	return nil
}
