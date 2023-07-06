package switchEntranceGameserver

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchEntranceGameserverGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwitchEntranceGameserverGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchEntranceGameserverGetLogic {
	return &SwitchEntranceGameserverGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwitchEntranceGameserverGetLogic) SwitchEntranceGameserverGet(req *types.GetSwitchEntranceGameserverReq) (resp *types.ListSwitchEntranceGameserverData, err error) {
	get, err := l.svcCtx.YunWeiRpc.SwitchEntranceGameserverGet(l.ctx, &yunweiclient.GetSwitchEntranceGameserverReq{SwitchEntranceGameserverId: req.SwitchEntranceGameserverId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListSwitchEntranceGameserverData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
