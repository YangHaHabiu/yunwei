package insideServer

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideServerGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideServerGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideServerGetLogic {
	return &InsideServerGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideServerGetLogic) InsideServerGet(req *types.GetInsideServerReq) (resp *types.ListInsideServerData, err error) {
	get, err := l.svcCtx.IntranetRpc.InsideServerGet(l.ctx, &intranetclient.GetInsideServerReq{InsideServerId: req.InsideServerId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListInsideServerData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
