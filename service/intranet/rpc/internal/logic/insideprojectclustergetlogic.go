package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProjectClusterGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideProjectClusterGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProjectClusterGetLogic {
	return &InsideProjectClusterGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideProjectClusterGetLogic) InsideProjectClusterGet(in *intranetclient.GetInsideProjectClusterReq) (*intranetclient.ListInsideProjectClusterData, error) {
	one, err := l.svcCtx.InsideProjectClusterModel.FindOne(l.ctx, in.InsideProjectClusterId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp intranetclient.ListInsideProjectClusterData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
