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

type InsideProjectClusterUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideProjectClusterUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProjectClusterUpdateLogic {
	return &InsideProjectClusterUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideProjectClusterUpdateLogic) InsideProjectClusterUpdate(in *intranetclient.UpdateInsideProjectClusterReq) (*intranetclient.InsideProjectClusterCommonResp, error) {
	var tmp model.InsideProjectCluster
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝更新数据失败，原因：" + err.Error())
	}
	err = l.svcCtx.InsideProjectClusterModel.Update(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideProjectClusterCommonResp{}, nil
}
