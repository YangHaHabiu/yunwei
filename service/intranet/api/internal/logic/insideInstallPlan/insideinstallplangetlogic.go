package insideInstallPlan

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideInstallPlanGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideInstallPlanGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideInstallPlanGetLogic {
	return &InsideInstallPlanGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideInstallPlanGetLogic) InsideInstallPlanGet(req *types.GetInsideInstallPlanReq) (resp *types.ListInsideInstallPlanData, err error) {
	get, err := l.svcCtx.IntranetRpc.InsideInstallPlanGet(l.ctx, &intranetclient.GetInsideInstallPlanReq{InsideInstallPlanId: req.InsideInstallPlanId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListInsideInstallPlanData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
