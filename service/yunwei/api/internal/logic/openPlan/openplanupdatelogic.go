package openPlan

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

type OpenPlanUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenPlanUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanUpdateLogic {
	return &OpenPlanUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenPlanUpdateLogic) OpenPlanUpdate(req *types.UpdateOpenPlanReq) error {
	var tmp yunwei.OpenPlanCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("更新拷贝数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.OpenPlanUpdate(l.ctx, &yunweiclient.UpdateOpenPlanReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
