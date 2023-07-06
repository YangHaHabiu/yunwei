package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProxyHostGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideProxyHostGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProxyHostGetLogic {
	return &InsideProxyHostGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideProxyHostGetLogic) InsideProxyHostGet(in *intranetclient.GetInsideProxyHostReq) (*intranetclient.ListInsideProxyHostData, error) {
	one, err := l.svcCtx.InsideProxyHostModel.FindOne(l.ctx, in.InsideProxyHostId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp intranetclient.ListInsideProxyHostData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
