package asset

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
	"ywadmin-v3/service/yunwei/rpc/yunwei"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"
)

type AssetAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetAddLogic {
	return &AssetAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetAddLogic) AssetAdd(req *types.AddAssetReq) error {

	var tmp []*yunwei.AssetDatas
	err2 := copier.Copy(&tmp, req.AssetData)
	if err2 != nil {
		return xerr.NewErrMsg("拷贝数据失败，原因：" + err2.Error())
	}
	_, err := l.svcCtx.YunWeiRpc.AssetAdd(l.ctx, &yunweiclient.AssetAddReq{Assetdata: tmp})
	if err != nil {
		return err
	}

	return nil
}
