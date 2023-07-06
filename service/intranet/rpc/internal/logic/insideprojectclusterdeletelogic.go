package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProjectClusterDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideProjectClusterDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProjectClusterDeleteLogic {
	return &InsideProjectClusterDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideProjectClusterDeleteLogic) InsideProjectClusterDelete(in *intranetclient.DeleteInsideProjectClusterReq) (*intranetclient.InsideProjectClusterCommonResp, error) {
	err := l.svcCtx.InsideProjectClusterModel.Delete(l.ctx, in.InsideProjectClusterId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideProjectClusterCommonResp{}, nil
}
