package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksLogsDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksLogsDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksLogsDeleteLogic {
	return &InsideTasksLogsDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksLogsDeleteLogic) InsideTasksLogsDelete(in *intranetclient.DeleteInsideTasksLogsReq) (*intranetclient.InsideTasksLogsCommonResp, error) {
	err := l.svcCtx.InsideTasksLogsModel.DeleteSoft(l.ctx, in.InsideTasksLogsId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideTasksLogsCommonResp{}, nil
}
