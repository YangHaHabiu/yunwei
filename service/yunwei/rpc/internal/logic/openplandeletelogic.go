package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenPlanDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenPlanDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenPlanDeleteLogic {
	return &OpenPlanDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenPlanDeleteLogic) OpenPlanDelete(in *yunweiclient.DeleteOpenPlanReq) (*yunweiclient.OpenPlanCommonResp, error) {
	tmps, err := l.svcCtx.OpenPlanModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败，原因：" + err.Error())
	}
	if tmps.InstallStatus == "1" && tmps.InitdbStatus == "1" {
		return nil, xerr.NewErrMsg("该计划已清理，已安装，禁止删除")
	}
	err = l.svcCtx.OpenPlanModel.DeleteSoft(l.ctx, in.Id)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &yunweiclient.OpenPlanCommonResp{}, nil
}
