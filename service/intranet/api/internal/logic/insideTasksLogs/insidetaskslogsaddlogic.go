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

type InsideTasksLogsAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideTasksLogsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksLogsAddLogic {
	return &InsideTasksLogsAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideTasksLogsAddLogic) InsideTasksLogsAdd(req *types.AddInsideTasksLogsReq) error {
	var tmp intranetclient.InsideTasksLogsCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.IntranetRpc.InsideTasksLogsAdd(l.ctx, &intranetclient.AddInsideTasksLogsReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
