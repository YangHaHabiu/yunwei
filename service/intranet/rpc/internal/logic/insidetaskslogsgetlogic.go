package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksLogsGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksLogsGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksLogsGetLogic {
	return &InsideTasksLogsGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksLogsGetLogic) InsideTasksLogsGet(in *intranetclient.GetInsideTasksLogsReq) (*intranetclient.ListInsideTasksLogsData, error) {
	one, err := l.svcCtx.InsideTasksLogsModel.FindOne(l.ctx, in.InsideTasksLogsId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp intranetclient.ListInsideTasksLogsData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
