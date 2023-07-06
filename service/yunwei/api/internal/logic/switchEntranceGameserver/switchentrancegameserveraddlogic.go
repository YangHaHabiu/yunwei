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

type SwitchEntranceGameserverAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwitchEntranceGameserverAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchEntranceGameserverAddLogic {
	return &SwitchEntranceGameserverAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwitchEntranceGameserverAddLogic) SwitchEntranceGameserverAdd(req *types.AddSwitchEntranceGameserverReq) error {
	var tmp yunweiclient.SwitchEntranceGameserverCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.SwitchEntranceGameserverAdd(l.ctx, &yunweiclient.AddSwitchEntranceGameserverReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil

}
