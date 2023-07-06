package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideHostInfoDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideHostInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideHostInfoDeleteLogic {
	return &InsideHostInfoDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideHostInfoDeleteLogic) InsideHostInfoDelete(in *intranetclient.DeleteInsideHostInfoReq) (*intranetclient.InsideHostInfoCommonResp, error) {

	all, err := l.svcCtx.InsideProxyHostModel.FindAll(l.ctx, "host_id__=", in.InsideHostInfoId)
	if len(*all) != 0 {
		return nil, xerr.NewErrMsg("存在关联的项目，请先删除项目")
	}
	err = l.svcCtx.InsideHostInfoModel.DeleteSoft(l.ctx, in.InsideHostInfoId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideHostInfoCommonResp{}, nil
}
