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

type InsideHostInfoAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideHostInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideHostInfoAddLogic {
	return &InsideHostInfoAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideHostInfoAddLogic) InsideHostInfoAdd(req *types.AddInsideHostInfoReq) error {
	var tmp intranetclient.InsideHostInfoCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.IntranetRpc.InsideHostInfoAdd(l.ctx, &intranetclient.AddInsideHostInfoReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
