package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GameServerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGameServerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GameServerListLogic {
	return &GameServerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Hosts Rpc End
func (l *GameServerListLogic) GameServerList(in *yunweiclient.ListGameServerReq) (*yunweiclient.ListGameServerResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "server_status__in", in.ServerStatus,
		"project_id__in", in.ProjectIds,
		"new_platform_info__in", in.NewPlatformInfo,
		"outer_ip@inner_ip__or__regexp", in.Ip,
		"game_server_title__regexp", in.GameServerTitle,
		"open_time__xrange", in.OpenTime,
		"server_status__in", in.ServerStatusX,
	)

	list, err := l.svcCtx.PlatformModel.FindPageGameServerListByPage(l.ctx, in.Current, in.PageSize, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询服务器信息失败，原因：" + err.Error())
	}
	count, err := l.svcCtx.PlatformModel.CountGameServer(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("统计服务器信息失败，原因：" + err.Error())
	}
	var tmp []*yunweiclient.ListGameServerData
	err = copier.Copy(&tmp, list)
	if err != nil {
		return nil, xerr.NewErrMsg("复制服务器信息失败，原因：" + err.Error())
	}
	return &yunweiclient.ListGameServerResp{
		Rows:  tmp,
		Total: count,
	}, nil

}
