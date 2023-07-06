package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergePlanDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMergePlanDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanDeleteLogic {
	return &MergePlanDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MergePlanDeleteLogic) MergePlanDelete(in *yunweiclient.DeleteMergePlanReq) (*yunweiclient.MergePlanCommonResp, error) {
	one, _ := l.svcCtx.MergePlanModel.FindOne(l.ctx, in.Id)
	if one.MergeStatus != -1 {
		return nil, xerr.NewErrMsg("已合服不允许删除")
	}
	err := l.svcCtx.MergePlanModel.DeleteSoft(l.ctx, in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &yunweiclient.MergePlanCommonResp{}, nil
}
