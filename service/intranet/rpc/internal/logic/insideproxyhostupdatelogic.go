package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/model"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideProxyHostUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideProxyHostUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProxyHostUpdateLogic {
	return &InsideProxyHostUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideProxyHostUpdateLogic) InsideProxyHostUpdate(in *intranetclient.UpdateInsideProxyHostReq) (*intranetclient.InsideProxyHostCommonResp, error) {
	var tmp model.InsideProxyHost
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝更新数据失败，原因：" + err.Error())
	}
	err = l.svcCtx.InsideProxyHostModel.Update(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideProxyHostCommonResp{}, nil
}
