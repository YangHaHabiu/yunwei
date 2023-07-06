package maintainPlan

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunwei"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaintainPlanUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanUpdateLogic {
	return &MaintainPlanUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaintainPlanUpdateLogic) MaintainPlanUpdate(req *types.UpdateMaintainPlanReq) error {
	var tmp yunwei.MaintainPlanCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("更新拷贝数据失败，原因：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.MaintainPlanUpdate(l.ctx, &yunweiclient.UpdateMaintainPlanReq{One: &tmp})
	if err != nil {
		return err
	}

	return nil
}
