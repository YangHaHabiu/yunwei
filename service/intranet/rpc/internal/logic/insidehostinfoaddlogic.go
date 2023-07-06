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

type InsideHostInfoAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideHostInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideHostInfoAddLogic {
	return &InsideHostInfoAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsideHostInfo Rpc Start
func (l *InsideHostInfoAddLogic) InsideHostInfoAdd(in *intranetclient.AddInsideHostInfoReq) (*intranetclient.InsideHostInfoCommonResp, error) {
	var tmp model.InsideHostInfo
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}

	_, err = l.svcCtx.InsideHostInfoModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	return &intranetclient.InsideHostInfoCommonResp{}, nil
}
