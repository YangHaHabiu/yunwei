package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideTasksPidGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksPidGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksPidGetLogic {
	return &InsideTasksPidGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksPidGetLogic) InsideTasksPidGet(in *intranetclient.GetInsideTasksPidReq) (*intranetclient.ListInsideTasksPidData, error) {
	one, err := l.svcCtx.InsideTasksPidModel.FindOne(l.ctx, in.InsideTasksPidId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp intranetclient.ListInsideTasksPidData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
