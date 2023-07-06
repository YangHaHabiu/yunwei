package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssetAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetAddLogic {
	return &AssetAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// asset rpc start
func (l *AssetAddLogic) AssetAdd(in *yunweiclient.AssetAddReq) (*yunweiclient.AssetAddResp, error) {
	err := l.svcCtx.AssetModel.BulkInserter(l.ctx, in)
	if err != nil {
		return nil, xerr.NewErrMsg("批量插入资产失败，原因：" + err.Error())
	}
	return &yunweiclient.AssetAddResp{}, nil
}
