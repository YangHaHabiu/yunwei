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

type InsideServerAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideServerAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideServerAddLogic {
	return &InsideServerAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsideServer Rpc Start
func (l *InsideServerAddLogic) InsideServerAdd(in *intranetclient.AddInsideServerReq) (*intranetclient.InsideServerCommonResp, error) {
	var tmp model.InsideServer
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.InsideServerModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideServerCommonResp{}, nil
}
