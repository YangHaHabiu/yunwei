package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoOpengameRuleGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAutoOpengameRuleGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoOpengameRuleGetLogic {
	return &AutoOpengameRuleGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AutoOpengameRuleGetLogic) AutoOpengameRuleGet(in *yunweiclient.GetAutoOpengameRuleReq) (*yunweiclient.ListAutoOpengameRuleData, error) {
	one, err := l.svcCtx.AutoOpengameRuleModel.FindOne(l.ctx, in.AutoOpengameRuleId)
	if err != nil {
		return nil, xerr.NewErrMsg("查询单条数据失败")
	}
	var tmp yunweiclient.ListAutoOpengameRuleData
	err = copier.Copy(&tmp, one)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝查询单条数据失败，原因：" + err.Error())
	}

	return &tmp, nil
}
