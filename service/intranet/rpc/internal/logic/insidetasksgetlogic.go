package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksGetLogic {
	return &InsideTasksGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksGetLogic) InsideTasksGet(in *intranetclient.GetInsideTasksReq) (*intranetclient.ListInsideTasksData, error) {
	one, err := l.svcCtx.InsideTasksModel.FindOne(l.ctx, in.InsideTasksId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp intranetclient.ListInsideTasksData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
