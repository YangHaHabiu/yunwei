package gameServer

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/api/internal/logic/common"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GameServerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameServerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GameServerListLogic {
	return &GameServerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GameServerListLogic) GameServerList(req *types.ListGameServerReq) (resp *types.ListGameServerResp, err error) {
	//个人项目值及列表
	projectIds, projectList, err := common.GetProjectStrAndList(l.svcCtx, l.ctx, req.ProjectIds)
	if err != nil {
		return nil, err
	}
	//获取平台
	platformList, err := common.GetPlatform(l.svcCtx, l.ctx, projectIds, "", req.PlatformType)
	if err != nil {
		return nil, err
	}
	list, err := l.svcCtx.YunWeiRpc.GameServerList(l.ctx, &yunweiclient.ListGameServerReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		ServerStatus:    req.ServerStatus,
		NewPlatformInfo: req.NewPlatformInfo,
		Ip:              req.Ip,
		GameServerTitle: req.GameServerTitle,
		OpenTime:        req.OpenTime,
		ProjectIds:      projectIds,
		ServerStatusX:   req.ServerStatusX,
	})

	if err != nil {
		return nil, err
	}
	tmp := make([]*types.ListGameServerData, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制游戏服信息出错，原因：" + err.Error())
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
			Label:    "平台",
			Value:    "newPlatformInfo",
			Types:    "select",
			Children: platformList,
		},
		{
			Label: "服名",
			Value: "gameServerTitle",
			Types: "input",
		},
		{
			Label: "IP",
			Value: "ip",
			Types: "input",
		},
		{
			Label: "服状态",
			Value: "serverStatusX",
			Types: "select",
			Children: []*types.FilterList{
				{
					Label: "收费服",
					Value: "1",
				}, {
					Label: "合服",
					Value: "2",
				},
			},
		},
	}

	resp = new(types.ListGameServerResp)
	resp.Rows = tmp
	resp.Filter = filterList
	resp.Total = list.Total

	return
}
