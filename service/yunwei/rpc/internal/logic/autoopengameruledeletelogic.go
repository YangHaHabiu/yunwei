package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoOpengameRuleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAutoOpengameRuleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoOpengameRuleDeleteLogic {
	return &AutoOpengameRuleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AutoOpengameRuleDeleteLogic) AutoOpengameRuleDelete(in *yunweiclient.DeleteAutoOpengameRuleReq) (*yunweiclient.AutoOpengameRuleCommonResp, error) {
	err := l.svcCtx.AutoOpengameRuleModel.DeleteSoft(l.ctx, in.AutoOpengameRuleId)
	if err != nil {
		return nil, xerr.NewErrMsg("删除信息失败，原因：" + err.Error())
	}
	return &yunweiclient.AutoOpengameRuleCommonResp{}, nil
}
