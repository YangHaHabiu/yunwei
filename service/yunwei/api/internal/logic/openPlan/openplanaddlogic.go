package openPlan

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenPlanAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenPlanAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanAddLogic {
	return &OpenPlanAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenPlanAddLogic) OpenPlanAdd(req *types.AddOpenPlanReq) error {
	var tmp []*yunweiclient.OpenPlanCommon
	err := copier.Copy(&tmp, req.OpenPlanData)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}
	_, err = l.svcCtx.YunWeiRpc.OpenPlanAdd(l.ctx, &yunweiclient.AddOpenPlanReq{OpenPlatData: tmp})
	if err != nil {
		return err
	}
	return nil
}
