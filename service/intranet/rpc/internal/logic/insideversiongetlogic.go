package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideVersionGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideVersionGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideVersionGetLogic {
	return &InsideVersionGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideVersionGetLogic) InsideVersionGet(in *intranetclient.GetInsideVersionReq) (*intranetclient.ListInsideVersionData, error) {
	one, err := l.svcCtx.InsideVersionModel.FindOne(l.ctx, in.InsideVersionId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp intranetclient.ListInsideVersionData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
