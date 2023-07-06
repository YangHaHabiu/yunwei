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

type InsideInstallPlanUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideInstallPlanUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideInstallPlanUpdateLogic {
	return &InsideInstallPlanUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideInstallPlanUpdateLogic) InsideInstallPlanUpdate(in *intranetclient.UpdateInsideInstallPlanReq) (*intranetclient.InsideInstallPlanCommonResp, error) {
	var tmp model.InsideInstallPlan
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝更新数据失败，原因：" + err.Error())
	}
	err = l.svcCtx.InsideInstallPlanModel.Update(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}
	return &intranetclient.InsideInstallPlanCommonResp{}, nil
}
