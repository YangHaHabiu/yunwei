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

type PlatformGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlatformGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformGetLogic {
	return &PlatformGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlatformGetLogic) PlatformGet(req *types.GetPlatformReq) (*types.GetPlatformResp, error) {

	get, err := l.svcCtx.YunWeiRpc.PlatformGet(l.ctx, &yunweiclient.GetPlatformReq{PlatformId: req.PlatformId})
	if err != nil {
		return nil, err
	}
	tmp := new(types.ListPlatformData)
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}

	return &types.GetPlatformResp{
		Row: tmp,
	}, nil
}
