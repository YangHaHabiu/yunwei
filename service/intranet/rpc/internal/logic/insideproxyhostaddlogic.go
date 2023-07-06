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

type InsideProxyHostAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideProxyHostAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideProxyHostAddLogic {
	return &InsideProxyHostAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsideProjectCluster Rpc End
func (l *InsideProxyHostAddLogic) InsideProxyHostAdd(in *intranetclient.AddInsideProxyHostReq) (*intranetclient.InsideProxyHostCommonResp, error) {
	var tmp model.InsideProxyHost
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}

	_, err = l.svcCtx.InsideProxyHostModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("存在重复的项目，请检查")
	}
	return &intranetclient.InsideProxyHostCommonResp{}, nil
}
