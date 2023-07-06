package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideServerDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideServerDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideServerDeleteLogic {
	return &InsideServerDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideServerDeleteLogic) InsideServerDelete(in *intranetclient.DeleteInsideServerReq) (*intranetclient.InsideServerCommonResp, error) {
	err := l.svcCtx.InsideServerModel.DeleteSoft(l.ctx, in.InsideServerId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideServerCommonResp{}, nil
}
