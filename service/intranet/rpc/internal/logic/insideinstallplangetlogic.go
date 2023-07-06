package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideInstallPlanGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideInstallPlanGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideInstallPlanGetLogic {
	return &InsideInstallPlanGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideInstallPlanGetLogic) InsideInstallPlanGet(in *intranetclient.GetInsideInstallPlanReq) (*intranetclient.ListInsideInstallPlanData, error) {
	one, err := l.svcCtx.InsideInstallPlanModel.FindOne(l.ctx, in.InsideInstallPlanId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp intranetclient.ListInsideInstallPlanData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
