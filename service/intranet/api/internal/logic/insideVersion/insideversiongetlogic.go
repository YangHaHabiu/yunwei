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

type InsideVersionGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideVersionGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideVersionGetLogic {
	return &InsideVersionGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideVersionGetLogic) InsideVersionGet(req *types.GetInsideVersionReq) (resp *types.ListInsideVersionData, err error) {
	get, err := l.svcCtx.IntranetRpc.InsideVersionGet(l.ctx, &intranetclient.GetInsideVersionReq{InsideVersionId: req.InsideVersionId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListInsideVersionData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
