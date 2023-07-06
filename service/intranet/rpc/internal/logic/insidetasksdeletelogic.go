package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksDeleteLogic {
	return &InsideTasksDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksDeleteLogic) InsideTasksDelete(in *intranetclient.DeleteInsideTasksReq) (*intranetclient.InsideTasksCommonResp, error) {
	err := l.svcCtx.InsideTasksModel.DeleteSoft(l.ctx, in.InsideTasksId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideTasksCommonResp{}, nil
}
