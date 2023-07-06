package asset

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetGetLogic {
	return &AssetGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetGetLogic) AssetGet(req *types.GetAssetReq) (resp *types.AddAssetResp, err error) {
	tmp := new(types.ListAssetData)
	list, err := l.svcCtx.YunWeiRpc.AssetList(l.ctx, &yunweiclient.AssetListReq{
		Current:  0,
		PageSize: 0,
		AssetId:  req.AssetId,
	})

	if err != nil || len(list.List) != 1 {
		return nil, xerr.NewErrMsg("获取单条资产数据失败")
	}

	err = copier.Copy(&tmp, list.List[0])
	if err != nil {
		return nil, xerr.NewErrMsg("复制单条资产数据失败，原因" + err.Error())
	}
	resp = new(types.AddAssetResp)
	resp.Row = tmp
	return
}
