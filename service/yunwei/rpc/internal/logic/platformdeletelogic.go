package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlatformDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlatformDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformDeleteLogic {
	return &PlatformDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlatformDeleteLogic) PlatformDelete(in *yunweiclient.DeletePlatformReq) (*yunweiclient.PlatformCommonResp, error) {

	err := l.svcCtx.PlatformModel.DeleteSoft(l.ctx, in.PlatformId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &yunweiclient.PlatformCommonResp{}, nil
}
