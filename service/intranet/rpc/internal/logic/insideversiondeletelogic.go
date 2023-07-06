package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideVersionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideVersionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideVersionDeleteLogic {
	return &InsideVersionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideVersionDeleteLogic) InsideVersionDelete(in *intranetclient.DeleteInsideVersionReq) (*intranetclient.InsideVersionCommonResp, error) {
	err := l.svcCtx.InsideVersionModel.DeleteSoft(l.ctx, in.InsideVersionId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideVersionCommonResp{}, nil
}
