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

type AutoOpengameRuleGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAutoOpengameRuleGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoOpengameRuleGetLogic {
	return &AutoOpengameRuleGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AutoOpengameRuleGetLogic) AutoOpengameRuleGet(req *types.GetAutoOpengameRuleReq) (resp *types.ListAutoOpengameRuleData, err error) {
	get, err := l.svcCtx.YunWeiRpc.AutoOpengameRuleGet(l.ctx, &yunweiclient.GetAutoOpengameRuleReq{AutoOpengameRuleId: req.AutoOpengameRuleId})
	if err != nil {
		return nil, err
	}
	var tmp types.ListAutoOpengameRuleData
	err = copier.Copy(&tmp, get)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝单条数据失败，原因：" + err.Error())
	}
	return &tmp, nil
}
