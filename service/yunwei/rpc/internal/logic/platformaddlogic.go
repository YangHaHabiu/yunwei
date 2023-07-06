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

type PlatformAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlatformAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformAddLogic {
	return &PlatformAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Platform rpc start
func (l *PlatformAddLogic) PlatformAdd(in *yunweiclient.AddPlatformReq) (*yunweiclient.PlatformCommonResp, error) {
	var tmp model.Platform
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.PlatformModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}
	return &yunweiclient.PlatformCommonResp{}, nil
}
