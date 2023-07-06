package platform

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlatformAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlatformAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformAddLogic {
	return &PlatformAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlatformAddLogic) PlatformAdd(req *types.AddPlatformReq) error {
	var tmp yunweiclient.PlatformCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.PlatformAdd(l.ctx, &yunweiclient.AddPlatformReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
