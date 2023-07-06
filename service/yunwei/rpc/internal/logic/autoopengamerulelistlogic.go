package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/model"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoOpengameRuleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAutoOpengameRuleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoOpengameRuleListLogic {
	return &AutoOpengameRuleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AutoOpengameRuleListLogic) AutoOpengameRuleList(in *yunweiclient.ListAutoOpengameRuleReq) (*yunweiclient.ListAutoOpengameRuleResp, error) {
	var (
		count int64
		list  []*yunweiclient.ListAutoOpengameRuleData
		all   *[]model.AutoOpengameRuleNew
		err   error
	)

	filters := make([]interface{}, 0)
	filters = append(filters, "project_id__in", in.ProjectIds)
	count, _ = l.svcCtx.AutoOpengameRuleModel.Count(l.ctx, filters...)
	if in.Current == 0 && in.PageSize == 0 {
		all, err = l.svcCtx.AutoOpengameRuleModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.AutoOpengameRuleModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	}
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}

	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}
	return &yunweiclient.ListAutoOpengameRuleResp{
		Rows:  list,
		Total: count,
	}, nil
}
