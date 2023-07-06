package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/model"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProjectClusterAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideProjectClusterAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProjectClusterAddLogic {
	return &InsideProjectClusterAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsideOperation Rpc End
func (l *InsideProjectClusterAddLogic) InsideProjectClusterAdd(in *intranetclient.AddInsideProjectClusterReq) (*intranetclient.InsideProjectClusterCommonResp, error) {
	var tmp model.InsideProjectCluster
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.InsideProjectClusterModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideProjectClusterCommonResp{}, nil
}
