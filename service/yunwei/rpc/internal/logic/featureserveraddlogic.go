package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeatureServerAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeatureServerAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeatureServerAddLogic {
	return &FeatureServerAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// feature_server rpc start
func (l *FeatureServerAddLogic) FeatureServerAdd(in *yunweiclient.AddFeatureServerReq) (*yunweiclient.FeatureServerCommonResp, error) {
	_, err := l.svcCtx.FeatureServerInfoModel.Insert(l.ctx, &model.FeatureServerInfo{
		FeatureServerId:   in.One.FeatureServerId,
		ProjectId:         in.One.ProjectId,
		FeatureServerInfo: in.One.FeatureServerInfo,
		Remark:            in.One.Remark,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}
	return &yunweiclient.FeatureServerCommonResp{}, nil
}
