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

type InsideOperationAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideOperationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideOperationAddLogic {
	return &InsideOperationAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideOperationAddLogic) InsideOperationAdd(req *types.AddInsideOperationReq) error {
	var tmp intranetclient.InsideOperationCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.IntranetRpc.InsideOperationAdd(l.ctx, &intranetclient.AddInsideOperationReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
