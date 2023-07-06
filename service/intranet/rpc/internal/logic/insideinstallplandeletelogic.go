package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/intranet/rpc/internal/svc"
	"ywadmin-v3/service/intranet/rpc/intranetclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsideInstallPlanDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideInstallPlanDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideInstallPlanDeleteLogic {
	return &InsideInstallPlanDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsideInstallPlanDeleteLogic) InsideInstallPlanDelete(in *intranetclient.DeleteInsideInstallPlanReq) (*intranetclient.InsideInstallPlanCommonResp, error) {
	err := l.svcCtx.InsideInstallPlanModel.DeleteSoft(l.ctx, in.InsideInstallPlanId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideInstallPlanCommonResp{}, nil
}
