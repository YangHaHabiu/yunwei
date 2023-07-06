package logic

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net/url"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetRecycleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssetRecycleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetRecycleDeleteLogic {
	return &AssetRecycleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssetRecycleDeleteLogic) AssetRecycleDelete(in *yunweiclient.AssetRecycleDeleteReq) (*yunweiclient.AssetRecycleDeleteResp, error) {
	var uname string
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		uname = md.Get("nickName")[0]
		uname, _ = url.QueryUnescape(uname)

	}
	err := l.svcCtx.AssetModel.RecycleDeleteSoft(l.ctx, in.AssetId, in.RecycleType, uname)
	if err != nil {
		return nil, xerr.NewErrMsg("删除资产失败，原因：" + err.Error())
	}

	return &yunweiclient.AssetRecycleDeleteResp{}, nil
}
