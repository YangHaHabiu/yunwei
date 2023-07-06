package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeatureServerDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeatureServerDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeatureServerDeleteLogic {
	return &FeatureServerDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeatureServerDeleteLogic) FeatureServerDelete(in *yunweiclient.DeleteFeatureServerReq) (*yunweiclient.FeatureServerCommonResp, error) {
	err := l.svcCtx.FeatureServerInfoModel.DeleteSoft(l.ctx, in.FeatureServerId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &yunweiclient.FeatureServerCommonResp{}, nil
}
