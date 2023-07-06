package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideHostInfoGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideHostInfoGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideHostInfoGetLogic {
	return &InsideHostInfoGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideHostInfoGetLogic) InsideHostInfoGet(in *intranetclient.GetInsideHostInfoReq) (*intranetclient.ListInsideHostInfoData, error) {
	one, err := l.svcCtx.InsideHostInfoModel.FindOne(l.ctx, in.InsideHostInfoId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp intranetclient.ListInsideHostInfoData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
