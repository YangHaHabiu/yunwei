package maintainPlan

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainGetClusterInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaintainGetClusterInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainGetClusterInfoLogic {
	return &MaintainGetClusterInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaintainGetClusterInfoLogic) MaintainGetClusterInfo(req *types.MaintainGetClusterInfoReq) (resp *types.MaintainGetClusterInfoResp, err error) {
	tmp := make([]*types.MaintainGetClusterInfoData, 0)
	list, err := l.svcCtx.YunWeiRpc.MaintainGetClusterInfo(l.ctx, &yunweiclient.MaintainGetClusterInfoReq{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&tmp, list.ClusterInfoData)
	if err != nil {
		return nil, err
	}
	resp = new(types.MaintainGetClusterInfoResp)
	resp.ClusterInfoData = tmp
	return
}
