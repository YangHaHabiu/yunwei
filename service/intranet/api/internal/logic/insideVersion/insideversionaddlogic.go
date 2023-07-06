package insideVersion

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideVersionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideVersionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideVersionAddLogic {
	return &InsideVersionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideVersionAddLogic) InsideVersionAdd(req *types.AddInsideVersionReq) error {
	var tmp intranetclient.InsideVersionCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.IntranetRpc.InsideVersionAdd(l.ctx, &intranetclient.AddInsideVersionReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
