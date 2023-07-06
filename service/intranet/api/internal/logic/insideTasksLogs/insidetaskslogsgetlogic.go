package insideTasksLogs

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksLogsGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideTasksLogsGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksLogsGetLogic {
	return &InsideTasksLogsGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideTasksLogsGetLogic) InsideTasksLogsGet(req *types.GetInsideTasksLogsReq) (resp *types.ListInsideTasksLogsData, err error) {
	get, err := l.svcCtx.IntranetRpc.InsideTasksLogsGet(l.ctx, &intranetclient.GetInsideTasksLogsReq{InsideTasksLogsId: req.InsideTasksLogsId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListInsideTasksLogsData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
