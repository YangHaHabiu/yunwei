package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaintainGetClusterInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMaintainGetClusterInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaintainGetClusterInfoLogic {
	return &MaintainGetClusterInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MaintainGetClusterInfoLogic) MaintainGetClusterInfo(in *yunweiclient.MaintainGetClusterInfoReq) (*yunweiclient.MaintainGetClusterInfoResp, error) {
	info, err := l.svcCtx.MaintainPlanModel.FindAllClusterInfo(l.ctx, "view_project_id__=", in.ProjectId,
		"view_cluster_label_id__in", in.LabelIds)
	if err != nil {
		return nil, xerr.NewErrMsg("查询项目所属集群信息失败，原因：" + err.Error())
	}

	var tmp []*yunweiclient.MaintainGetClusterInfoData
	err = copier.Copy(&tmp, info)
	if err != nil {
		return nil, xerr.NewErrMsg("复制项目所属集群信息失败，原因：" + err.Error())
	}

	return &yunweiclient.MaintainGetClusterInfoResp{
		ClusterInfoData: tmp,
	}, nil
}
