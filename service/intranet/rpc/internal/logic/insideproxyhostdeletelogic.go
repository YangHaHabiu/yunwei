package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProxyHostDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideProxyHostDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProxyHostDeleteLogic {
	return &InsideProxyHostDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideProxyHostDeleteLogic) InsideProxyHostDelete(in *intranetclient.DeleteInsideProxyHostReq) (*intranetclient.InsideProxyHostCommonResp, error) {
	err := l.svcCtx.InsideProxyHostModel.Delete(l.ctx, in.InsideProxyHostId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideProxyHostCommonResp{}, nil
}
