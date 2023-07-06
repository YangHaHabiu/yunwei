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

type InsideInstallPlanAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsideInstallPlanAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsideInstallPlanAddLogic {
	return &InsideInstallPlanAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsideFeature Rpc End
func (l *InsideInstallPlanAddLogic) InsideInstallPlanAdd(in *intranetclient.AddInsideInstallPlanReq) (*intranetclient.InsideInstallPlanCommonResp, error) {
	var tmp model.InsideInstallPlan
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝新增数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.InsideInstallPlanModel.Insert(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("插入信息失败，原因：" + err.Error())
	}

	return &intranetclient.InsideInstallPlanCommonResp{}, nil
}
