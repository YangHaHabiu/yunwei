package autoOpengameRule

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoOpengameRuleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAutoOpengameRuleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoOpengameRuleListLogic {
	return &AutoOpengameRuleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AutoOpengameRuleListLogic) AutoOpengameRuleList(req *types.ListAutoOpengameRuleReq) (resp *types.ListAutoOpengameRuleResp, err error) {
	tmp := make([]*types.ListAutoOpengameRuleData, 0)
	projectIds, _, _ := common.GetProjectStrAndList(l.svcCtx, l.ctx, "")

	list, err := l.svcCtx.YunWeiRpc.AutoOpengameRuleList(l.ctx, &yunweiclient.ListAutoOpengameRuleReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		ProjectIds: projectIds,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	return &types.ListAutoOpengameRuleResp{
		Rows:  tmp,
		Total: list.Total,
	}, nil
}
