package insideHostInfo

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideHostInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideHostInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideHostInfoListLogic {
	return &InsideHostInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideHostInfoListLogic) InsideHostInfoList(req *types.ListInsideHostInfoReq) (resp *types.ListInsideHostInfoResp, err error) {
	tmp := make([]*types.ListInsideHostInfoData, 0)
	list, err := l.svcCtx.IntranetRpc.InsideHostInfoList(l.ctx, &intranetclient.ListInsideHostInfoReq{
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

	return &types.ListInsideHostInfoResp{
		Rows:  tmp,
		Total: list.Total,
	}, nil
}
