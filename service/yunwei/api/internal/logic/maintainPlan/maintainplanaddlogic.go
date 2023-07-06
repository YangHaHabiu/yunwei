package maintainPlan

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaintainPlanAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanAddLogic {
	return &MaintainPlanAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaintainPlanAddLogic) MaintainPlanAdd(req *types.AddMaintainPlanReq) error {
	var tmp yunweiclient.MaintainPlanCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.MaintainPlanAdd(l.ctx, &yunweiclient.AddMaintainPlanReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
