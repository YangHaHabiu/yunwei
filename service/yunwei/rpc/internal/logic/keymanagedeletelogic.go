package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type KeyManageDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKeyManageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeyManageDeleteLogic {
	return &KeyManageDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KeyManageDeleteLogic) KeyManageDelete(in *yunweiclient.DeleteKeyManageReq) (*yunweiclient.KeyManageCommonResp, error) {
	err := l.svcCtx.KeyManageModel.DeleteSoft(l.ctx, in.KeyId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &yunweiclient.KeyManageCommonResp{}, nil
}
