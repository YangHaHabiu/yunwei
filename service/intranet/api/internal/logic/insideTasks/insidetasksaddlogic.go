package insideTasks

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideTasksAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksAddLogic {
	return &InsideTasksAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideTasksAddLogic) InsideTasksAdd(req *types.AddInsideTasksReq) error {
	var tmp intranetclient.InsideTasksCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.IntranetRpc.InsideTasksAdd(l.ctx, &intranetclient.AddInsideTasksReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
