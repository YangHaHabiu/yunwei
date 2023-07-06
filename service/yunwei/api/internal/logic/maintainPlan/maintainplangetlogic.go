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

type MaintainPlanGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaintainPlanGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanGetLogic {
	return &MaintainPlanGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaintainPlanGetLogic) MaintainPlanGet(req *types.GetMaintainPlanReq) (resp *types.ListMaintainPlanData, err error) {
	get, err := l.svcCtx.YunWeiRpc.MaintainPlanGet(l.ctx, &yunweiclient.GetMaintainPlanReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	var tmp types.ListMaintainPlanData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
