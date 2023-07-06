package insideProxyHost

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProxyHostAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideProxyHostAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProxyHostAddLogic {
	return &InsideProxyHostAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideProxyHostAddLogic) InsideProxyHostAdd(req *types.AddInsideProxyHostReq) error {
	var tmp intranetclient.InsideProxyHostCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.IntranetRpc.InsideProxyHostAdd(l.ctx, &intranetclient.AddInsideProxyHostReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
