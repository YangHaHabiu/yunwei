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

type InsideProjectClusterAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideProjectClusterAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProjectClusterAddLogic {
	return &InsideProjectClusterAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideProjectClusterAddLogic) InsideProjectClusterAdd(req *types.AddInsideProjectClusterReq) error {
	var tmp intranetclient.InsideProjectClusterCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.IntranetRpc.InsideProjectClusterAdd(l.ctx, &intranetclient.AddInsideProjectClusterReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
