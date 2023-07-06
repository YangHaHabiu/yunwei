package insideOperation

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideOperationGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideOperationGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideOperationGetLogic {
	return &InsideOperationGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideOperationGetLogic) InsideOperationGet(req *types.GetInsideOperationReq) (resp *types.ListInsideOperationData, err error) {
	get, err := l.svcCtx.IntranetRpc.InsideOperationGet(l.ctx, &intranetclient.GetInsideOperationReq{InsideOperationId: req.InsideOperationId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListInsideOperationData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
