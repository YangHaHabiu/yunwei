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

type OpenPlanBatchModifyOpenTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenPlanBatchModifyOpenTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanBatchModifyOpenTimeLogic {
	return &OpenPlanBatchModifyOpenTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenPlanBatchModifyOpenTimeLogic) OpenPlanBatchModifyOpenTime(req *types.BatchModifyOpenTimeReq) error {
	var tmp []*yunweiclient.OpenPlanBatchModifyOpenTimeData
	err := copier.Copy(&tmp, req.Data)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.OpenPlanBatchModifyOpenTime(l.ctx, &yunweiclient.OpenPlanBatchModifyOpenTimeReq{Data: tmp})
	if err != nil {
		return err
	}

	return nil
}
