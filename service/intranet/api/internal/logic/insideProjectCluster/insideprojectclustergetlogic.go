package insideProjectCluster

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProjectClusterGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideProjectClusterGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProjectClusterGetLogic {
	return &InsideProjectClusterGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideProjectClusterGetLogic) InsideProjectClusterGet(req *types.GetInsideProjectClusterReq) (resp *types.ListInsideProjectClusterData, err error) {
	get, err := l.svcCtx.IntranetRpc.InsideProjectClusterGet(l.ctx, &intranetclient.GetInsideProjectClusterReq{InsideProjectClusterId: req.InsideProjectClusterId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListInsideProjectClusterData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
