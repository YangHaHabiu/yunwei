package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenPlanBatchModifyOpenTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenPlanBatchModifyOpenTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanBatchModifyOpenTimeLogic {
	return &OpenPlanBatchModifyOpenTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenPlanBatchModifyOpenTimeLogic) OpenPlanBatchModifyOpenTime(in *yunweiclient.OpenPlanBatchModifyOpenTimeReq) (*yunweiclient.OpenPlanBatchModifyOpenTimeResp, error) {
	err := l.svcCtx.OpenPlanModel.TransactBatchUpdateTime(l.ctx, in.Data)
	if err != nil {
		return nil, xerr.NewErrMsg("批量修改开服时间失败，原因：" + err.Error())
	}
	return &yunweiclient.OpenPlanBatchModifyOpenTimeResp{}, nil
}
