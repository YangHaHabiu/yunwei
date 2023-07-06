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

type InsideTasksPidUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideTasksPidUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideTasksPidUpdateLogic {
	return &InsideTasksPidUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideTasksPidUpdateLogic) InsideTasksPidUpdate(in *intranetclient.UpdateInsideTasksPidReq) (*intranetclient.InsideTasksPidCommonResp, error) {
	var tmp model.InsideTasksPid
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝更新数据失败，原因：" + err.Error())
	}
	err = l.svcCtx.InsideTasksPidModel.Update(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideTasksPidCommonResp{}, nil
}
