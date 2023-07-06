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

type AssetUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssetUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetUpdateLogic {
	return &AssetUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssetUpdateLogic) AssetUpdate(in *yunweiclient.AssetUpdateReq) (*yunweiclient.AssetUpdateResp, error) {

	var uname string
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		uname = md.Get("nickName")[0]
		uname, _ = url.QueryUnescape(uname)
	}
	err := l.svcCtx.AssetModel.UpdateNews(l.ctx, uname, in, l.svcCtx.Config.Scripts.InitScriptPath)
	if err != nil {
		return nil, xerr.NewErrMsg("修改资产信息失败，原因：" + err.Error())
	}

	return &yunweiclient.AssetUpdateResp{}, nil
}
