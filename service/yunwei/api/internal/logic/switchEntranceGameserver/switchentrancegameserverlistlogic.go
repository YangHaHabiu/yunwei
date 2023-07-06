package switchEntranceGameserver

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchEntranceGameserverListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwitchEntranceGameserverListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchEntranceGameserverListLogic {
	return &SwitchEntranceGameserverListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwitchEntranceGameserverListLogic) SwitchEntranceGameserverList(req *types.ListSwitchEntranceGameserverReq) (resp *types.ListSwitchEntranceGameserverResp, err error) {
	projectIds, _, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, "")
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListSwitchEntranceGameserverData, 0)
	list, err := l.svcCtx.YunWeiRpc.SwitchEntranceGameserverList(l.ctx, &yunweiclient.ListSwitchEntranceGameserverReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		ProjectIds: projectIds,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	return &types.ListSwitchEntranceGameserverResp{
		Rows:  tmp,
		Total: list.Total,
	}, nil
}
