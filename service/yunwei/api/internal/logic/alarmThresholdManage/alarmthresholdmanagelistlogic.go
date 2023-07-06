package alarmThresholdManage

import (
	"context"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmThresholdManageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlarmThresholdManageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmThresholdManageListLogic {
	return &AlarmThresholdManageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlarmThresholdManageListLogic) AlarmThresholdManageList(req *types.ListAlarmThresholdManageReq) (resp *types.ListAlarmThresholdManageResp, err error) {
	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListAlarmThresholdManageData, 0)
	list, err := l.svcCtx.YunWeiRpc.AlarmThresholdManageList(l.ctx, &yunweiclient.ListAlarmThresholdManageReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		Ips:             req.Ips,
		Types:           req.Types,
		ProjectIds:      projectIds,
		GameServerAlias: req.GameServerAlias,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, err
	}

	tmTypes, err := common.GetThresholdManageTypesList(l.svcCtx, l.ctx)
	if err != nil {
		return nil, err

	}
	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label:    "项目",
			Value:    "projectIds",
			Types:    "select",
			Children: projectList,
		},
		{
			Label:    "阈值管理类型",
			Value:    "types",
			Types:    "select",
			Children: tmTypes,
		},
		{
			Label: "主机IP",
			Value: "ips",
			Types: "input",
		},
		{
			Label: "服别名",
			Value: "gameServerAlias",
			Types: "input",
		},
	}

	return &types.ListAlarmThresholdManageResp{
		Rows:   tmp,
		Total:  list.Total,
		Filter: filterList,
	}, nil
}
