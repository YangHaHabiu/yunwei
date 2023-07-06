package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchEntranceGameserverAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSwitchEntranceGameserverAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchEntranceGameserverAddLogic {
	return &SwitchEntranceGameserverAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SwitchEntranceGameserver Rpc Start
func (l *SwitchEntranceGameserverAddLogic) SwitchEntranceGameserverAdd(in *yunweiclient.AddSwitchEntranceGameserverReq) (*yunweiclient.SwitchEntranceGameserverCommonResp, error) {
	var tmp model.SwitchEntranceGameserver
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.SwitchEntranceGameserverModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}

	return &yunweiclient.SwitchEntranceGameserverCommonResp{}, nil
}
