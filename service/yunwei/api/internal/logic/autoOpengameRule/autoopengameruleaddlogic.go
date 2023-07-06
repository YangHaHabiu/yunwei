package autoOpengameRule

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoOpengameRuleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAutoOpengameRuleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoOpengameRuleAddLogic {
	return &AutoOpengameRuleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AutoOpengameRuleAddLogic) AutoOpengameRuleAdd(req *types.AddAutoOpengameRuleReq) error {
	var tmp yunweiclient.AutoOpengameRuleCommon
	err := copier.Copy(&tmp, req)
	if err != nil {
		return xerr.NewErrMsg("新增拷贝数据失败，原因1：" + err.Error())
	}

	_, err = l.svcCtx.YunWeiRpc.AutoOpengameRuleAdd(l.ctx, &yunweiclient.AddAutoOpengameRuleReq{One: &tmp})
	if err != nil {
		return err
	}
	return nil
}
