package logic

import (
	"context"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergePlanRangeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMergePlanRangeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergePlanRangeListLogic {
	return &MergePlanRangeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MergePlanRangeListLogic) MergePlanRangeList(in *yunweiclient.GetMergeListTreeByPlatIdAndProIdReq) (*yunweiclient.GetMergeListTreeByPlatIdAndProIdResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, " game_server.project_id__=", in.ProjectId, "game_server.platform_id__=", in.PlatformId)
	list, err := l.svcCtx.MergePlanModel.FindAllRangeList(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询维护范围失败，原因：" + err.Error())
	}
	var data string

	if len(*list) > 1 {
		return nil, xerr.NewErrMsg("查询维护范围失败")
	} else if len(*list) == 1 {
		data = (*list)[0].Data
	}
	return &yunweiclient.GetMergeListTreeByPlatIdAndProIdResp{
		Data: data,
	}, nil
}
