package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanRangeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMaintainPlanRangeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanRangeListLogic {
	return &MaintainPlanRangeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MaintainPlanRangeListLogic) MaintainPlanRangeList(in *yunweiclient.GetMaintainPlanListTreeByClsIdAndProIdReq) (*yunweiclient.GetMaintainPlanListTreeByClsIdAndProIdResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, " game_server.project_id__=", in.ProjectId,
		"view_args_platform_group__in", in.ClusterEn)
	list, err := l.svcCtx.MaintainPlanModel.FindAllRangeList(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询维护范围失败，原因：" + err.Error())
	}
	var data string

	if len(*list) > 1 {
		return nil, xerr.NewErrMsg("查询维护范围失败")
	} else if len(*list) == 1 {
		data = (*list)[0].List.String
	}

	return &yunweiclient.GetMaintainPlanListTreeByClsIdAndProIdResp{
		Data: data,
	}, err
}
