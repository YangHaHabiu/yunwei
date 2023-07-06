package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetBatchDistributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssetBatchDistributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetBatchDistributeLogic {
	return &AssetBatchDistributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssetBatchDistributeLogic) AssetBatchDistribute(in *yunweiclient.BatchDistributeReq) (*yunweiclient.BatchDistributeResp, error) {

	err := l.svcCtx.AssetModel.TransactBatchUpdate(l.ctx, in, l.svcCtx.Config.Scripts.InitScriptPath)
	if err != nil {
		return nil, xerr.NewErrMsg("批量修改资产信息失败，原因：" + err.Error())
	}

	return &yunweiclient.BatchDistributeResp{}, nil
}
