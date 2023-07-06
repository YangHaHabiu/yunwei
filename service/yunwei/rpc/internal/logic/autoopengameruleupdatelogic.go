package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoOpengameRuleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAutoOpengameRuleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoOpengameRuleUpdateLogic {
	return &AutoOpengameRuleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AutoOpengameRuleUpdateLogic) AutoOpengameRuleUpdate(in *yunweiclient.UpdateAutoOpengameRuleReq) (*yunweiclient.AutoOpengameRuleCommonResp, error) {
	var tmp model.AutoOpengameRule
	err := copier.Copy(&tmp, in.One)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝更新数据失败，原因：" + err.Error())
	}
	err = l.svcCtx.AutoOpengameRuleModel.Update(l.ctx, &tmp)
	if err != nil {
		return nil, xerr.NewErrMsg("更新信息失败，原因：" + err.Error())
	}
	return &yunweiclient.AutoOpengameRuleCommonResp{}, nil
}
