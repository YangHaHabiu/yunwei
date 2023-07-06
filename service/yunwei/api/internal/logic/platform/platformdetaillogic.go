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

type PlatformDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlatformDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformDetailLogic {
	return &PlatformDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlatformDetailLogic) PlatformDetail(req *types.DetailPlatformReq) (resp *types.DetailPlatformResp, err error) {

	list, err := l.svcCtx.YunWeiRpc.PlatformDetail(l.ctx, &yunweiclient.DetailPlatformReq{PlatformId: req.PlatformId})
	if err != nil {
		return nil, err
	}

	var tmp types.DetailPlatformResp
	err = copier.Copy(&tmp, list)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
