package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssetDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetDeleteLogic {
	return &AssetDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssetDeleteLogic) AssetDelete(in *yunweiclient.AssetDeleteReq) (*yunweiclient.AssetDeleteResp, error) {
	one, err := l.svcCtx.AssetModel.FindOne(l.ctx, in.AssetId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条资产失败，原因：" + err.Error())
	}
	if one.InitType == "1" {
		return nil, xerr.NewErrMsg("禁止删除已初始化资产，原因：" + err.Error())
	}
	err = l.svcCtx.AssetModel.DeleteSoft(l.ctx, in.AssetId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除资产失败，原因：" + err.Error())
	}
	return &yunweiclient.AssetDeleteResp{}, nil
}
