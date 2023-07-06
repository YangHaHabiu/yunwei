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

type SwitchEntranceGameserverListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSwitchEntranceGameserverListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchEntranceGameserverListLogic {
	return &SwitchEntranceGameserverListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SwitchEntranceGameserverListLogic) SwitchEntranceGameserverList(in *yunweiclient.ListSwitchEntranceGameserverReq) (*yunweiclient.ListSwitchEntranceGameserverResp, error) {
	var (
		count int64
		list  []*yunweiclient.ListSwitchEntranceGameserverData
		all   *[]model.SwitchEntranceGameserverNew
		err   error
	)

	filters := make([]interface{}, 0)
	filters = append(filters, "project_id__in", in.ProjectIds)
	count, _ = l.svcCtx.SwitchEntranceGameserverModel.Count(l.ctx, filters...)
	if in.Current == 0 && in.PageSize == 0 {
		all, err = l.svcCtx.SwitchEntranceGameserverModel.FindAll(l.ctx, filters...)
	} else {
		all, err = l.svcCtx.SwitchEntranceGameserverModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
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
	return &yunweiclient.ListSwitchEntranceGameserverResp{
		Rows:  list,
		Total: count,
	}, nil
}
