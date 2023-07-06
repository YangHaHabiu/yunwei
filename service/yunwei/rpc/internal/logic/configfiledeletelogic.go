package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeleteLogic {
	return &ConfigFileDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigFileDeleteLogic) ConfigFileDelete(in *yunweiclient.DeleteConfigFileReq) (*yunweiclient.ConfigFileCommonResp, error) {
	err := l.svcCtx.ConfigFileModel.Delete(l.ctx, in.ConfigFileId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &yunweiclient.ConfigFileCommonResp{}, nil
}
