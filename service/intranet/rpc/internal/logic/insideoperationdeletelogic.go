package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideOperationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideOperationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideOperationDeleteLogic {
	return &InsideOperationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideOperationDeleteLogic) InsideOperationDelete(in *intranetclient.DeleteInsideOperationReq) (*intranetclient.InsideOperationCommonResp, error) {
	err := l.svcCtx.InsideOperationModel.DeleteSoft(l.ctx, in.InsideOperationId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideOperationCommonResp{}, nil
}
