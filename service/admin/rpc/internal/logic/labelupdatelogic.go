package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LabelUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLabelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LabelUpdateLogic {
	return &LabelUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LabelUpdateLogic) LabelUpdate(in *adminclient.LabelUpdateReq) (*adminclient.LabelUpdateResp, error) {
	err := l.svcCtx.LabelModel.Update(l.ctx, &model.Label{
		LabelId:     in.LabelId,
		LabelName:   in.LabelName,
		LabelValues: in.LabelValues,
		LabelType:   in.LabelType,
		LabelRemark: in.LabelRemark,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}
	return &adminclient.LabelUpdateResp{}, nil
}
