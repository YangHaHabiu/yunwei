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

type InsideTasksLogsAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksLogsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksLogsAddLogic {
	return &InsideTasksLogsAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsideTasks Rpc End
func (l *InsideTasksLogsAddLogic) InsideTasksLogsAdd(in *intranetclient.AddInsideTasksLogsReq) (*intranetclient.InsideTasksLogsCommonResp, error) {
	var tmp model.InsideTasksLogs
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.InsideTasksLogsModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideTasksLogsCommonResp{}, nil
}
