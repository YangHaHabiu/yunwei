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

type AssetBatchDistributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetBatchDistributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetBatchDistributeLogic {
	return &AssetBatchDistributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetBatchDistributeLogic) AssetBatchDistribute(req *types.AssetBatchDistributeReq) error {

	var tmp yunweiclient.BatchDistributeReq
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("复制批量操作参数失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.AssetBatchDistribute(l.ctx, &tmp)
	if err != nil {
		return err
	}
	return nil
}
