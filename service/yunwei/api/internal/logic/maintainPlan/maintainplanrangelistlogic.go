package maintainPlan

import (
	"context"
	"encoding/json"
	"fmt"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainPlanRangeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaintainPlanRangeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainPlanRangeListLogic {
	return &MaintainPlanRangeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaintainPlanRangeListLogic) MaintainPlanRangeList(req *types.GetMaintainPlanListTreeByClsIdAndProIdReq) (*types.GetMaintainPlanListTreeByClsIdAndProIdResp, error) {

	list, err := l.svcCtx.YunWeiRpc.MaintainPlanRangeList(l.ctx, &yunweiclient.GetMaintainPlanListTreeByClsIdAndProIdReq{
		ProjectId: req.ProjectId,
		ClusterEn: req.ClusterCn,
	})
	if err != nil {
		return nil, err
	}
	data := make([]*types.VueTree, 0)
	if list.Data != "" {
		err = json.Unmarshal([]byte(list.Data), &data)
		if err != nil {
			return nil, xerr.NewErrMsg(fmt.Sprintf("解析维护计划树形json失败,失败：%v", err))
		}
	}

	return &types.GetMaintainPlanListTreeByClsIdAndProIdResp{
		MaRangeTreeData: data,
	}, err
}
