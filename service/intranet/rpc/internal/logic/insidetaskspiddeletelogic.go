package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksPidDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksPidDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksPidDeleteLogic {
	return &InsideTasksPidDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksPidDeleteLogic) InsideTasksPidDelete(in *intranetclient.DeleteInsideTasksPidReq) (*intranetclient.InsideTasksPidCommonResp, error) {
	err := l.svcCtx.InsideTasksPidModel.Delete(l.ctx, in.InsideTasksPidId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideTasksPidCommonResp{}, nil
}
