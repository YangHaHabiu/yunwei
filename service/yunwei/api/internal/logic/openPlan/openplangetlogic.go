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

type OpenPlanGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenPlanGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanGetLogic {
	return &OpenPlanGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenPlanGetLogic) OpenPlanGet(req *types.GetOpenPlanReq) (resp *types.ListOpenPlanData, err error) {
	get, err := l.svcCtx.YunWeiRpc.OpenPlanGet(l.ctx, &yunweiclient.GetOpenPlanReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	var tmp types.ListOpenPlanData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
