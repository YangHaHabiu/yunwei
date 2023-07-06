package insideVersion

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranet"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideVersionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideVersionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideVersionUpdateLogic {
	return &InsideVersionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideVersionUpdateLogic) InsideVersionUpdate(req *types.UpdateInsideVersionReq) error {
	var tmp intranet.InsideVersionCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("更新拷贝数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.IntranetRpc.InsideVersionUpdate(l.ctx, &intranetclient.UpdateInsideVersionReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
