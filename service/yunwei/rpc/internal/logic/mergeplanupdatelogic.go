package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergePlanUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMergePlanUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanUpdateLogic {
	return &MergePlanUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MergePlanUpdateLogic) MergePlanUpdate(in *yunweiclient.UpdateMergePlanReq) (*yunweiclient.MergePlanCommonResp, error) {

	err := l.svcCtx.MergePlanModel.Update(l.ctx, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}
	return &yunweiclient.MergePlanCommonResp{}, nil
}
