package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeatureServerUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeatureServerUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeatureServerUpdateLogic {
	return &FeatureServerUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeatureServerUpdateLogic) FeatureServerUpdate(in *yunweiclient.UpdateFeatureServerReq) (*yunweiclient.FeatureServerCommonResp, error) {
	err := l.svcCtx.FeatureServerInfoModel.Update(l.ctx, &model.FeatureServerInfo{
		FeatureServerId:   in.One.FeatureServerId,
		ProjectId:         in.One.ProjectId,
		FeatureServerInfo: in.One.FeatureServerInfo,
		Remark:            in.One.Remark,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}
	return &yunweiclient.FeatureServerCommonResp{}, nil
}
