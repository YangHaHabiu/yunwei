package insideHostInfo

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideHostInfoGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideHostInfoGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideHostInfoGetLogic {
	return &InsideHostInfoGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideHostInfoGetLogic) InsideHostInfoGet(req *types.GetInsideHostInfoReq) (resp *types.ListInsideHostInfoData, err error) {
	get, err := l.svcCtx.IntranetRpc.InsideHostInfoGet(l.ctx, &intranetclient.GetInsideHostInfoReq{InsideHostInfoId: req.InsideHostInfoId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListInsideHostInfoData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
