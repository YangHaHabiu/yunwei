package taskQueue

import (
	"context"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WaitInstallListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWaitInstallListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WaitInstallListLogic {
	return &WaitInstallListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WaitInstallListLogic) WaitInstallList(req *types.ListWaitInstallReq) (resp *types.ListWaitInstallResp, err error) {
	// todo: add your logic here and delete this line

	return
}
