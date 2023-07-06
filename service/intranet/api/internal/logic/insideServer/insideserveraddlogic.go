package insideServer

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideServerAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideServerAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideServerAddLogic {
	return &InsideServerAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideServerAddLogic) InsideServerAdd(req *types.AddInsideServerReq) error {
	var tmp intranetclient.InsideServerCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.IntranetRpc.InsideServerAdd(l.ctx, &intranetclient.AddInsideServerReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
