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

type InsideInstallPlanAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsideInstallPlanAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideInstallPlanAddLogic {
	return &InsideInstallPlanAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsideInstallPlanAddLogic) InsideInstallPlanAdd(req *types.AddInsideInstallPlanReq) error {
	var tmp intranetclient.InsideInstallPlanCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.IntranetRpc.InsideInstallPlanAdd(l.ctx, &intranetclient.AddInsideInstallPlanReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
