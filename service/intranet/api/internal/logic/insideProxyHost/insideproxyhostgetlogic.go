package insideProxyHost

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProxyHostGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideProxyHostGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProxyHostGetLogic {
	return &InsideProxyHostGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideProxyHostGetLogic) InsideProxyHostGet(req *types.GetInsideProxyHostReq) (resp *types.ListInsideProxyHostData, err error) {
	get, err := l.svcCtx.IntranetRpc.InsideProxyHostGet(l.ctx, &intranetclient.GetInsideProxyHostReq{InsideProxyHostId: req.InsideProxyHostId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListInsideProxyHostData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
