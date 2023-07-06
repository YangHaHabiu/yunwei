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

type InsideTasksGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideTasksGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksGetLogic {
	return &InsideTasksGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideTasksGetLogic) InsideTasksGet(req *types.GetInsideTasksReq) (resp *types.ListInsideTasksData, err error) {
	get, err := l.svcCtx.IntranetRpc.InsideTasksGet(l.ctx, &intranetclient.GetInsideTasksReq{InsideTasksId: req.InsideTasksId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListInsideTasksData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
