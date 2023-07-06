package insideProxyHost

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProxyHostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideProxyHostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProxyHostListLogic {
	return &InsideProxyHostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideProxyHostListLogic) InsideProxyHostList(req *types.ListInsideProxyHostReq) (resp *types.ListInsideProxyHostResp, err error) {
	tmp := make([]*types.ListInsideProxyHostData, 0)
	list, err := l.svcCtx.IntranetRpc.InsideProxyHostList(l.ctx, &intranetclient.ListInsideProxyHostReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	return &types.ListInsideProxyHostResp{
		Rows:  tmp,
		Total: list.Total,
	}, nil
}
