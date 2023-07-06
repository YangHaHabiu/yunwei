package asset

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetUpdateLogic {
	return &AssetUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetUpdateLogic) AssetUpdate(req *types.UpdateAssetReq) error {
	var tmp yunweiclient.AssetDatas
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("复制修改资产参数失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.AssetUpdate(l.ctx, &yunweiclient.AssetUpdateReq{
		Asset: &tmp,
	})
	if err != nil {
		return err
	}

	return nil
}
