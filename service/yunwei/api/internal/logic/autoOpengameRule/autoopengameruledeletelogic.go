package autoOpengameRule

import (
	"context"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoOpengameRuleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAutoOpengameRuleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoOpengameRuleDeleteLogic {
	return &AutoOpengameRuleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AutoOpengameRuleDeleteLogic) AutoOpengameRuleDelete(req *types.DeleteAutoOpengameRuleReq) error {
	_, err := l.svcCtx.YunWeiRpc.AutoOpengameRuleDelete(l.ctx, &yunweiclient.DeleteAutoOpengameRuleReq{AutoOpengameRuleId: req.AutoOpengameRuleId})
	if err != nil {
		return err
	}
	return nil
}
