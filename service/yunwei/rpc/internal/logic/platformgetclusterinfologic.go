package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlatformGetClusterInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlatformGetClusterInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformGetClusterInfoLogic {
	return &PlatformGetClusterInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlatformGetClusterInfoLogic) PlatformGetClusterInfo(in *yunweiclient.GetClusterByPlatformReq) (*yunweiclient.GetClusterByPlatformResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "project_id__=", in.ProjectId,
		"platform_en__in", in.PlatformEns)
	ens, err := l.svcCtx.PlatformModel.FindClusterByPlatformEns(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询平台集群信息失败，原因：" + err.Error())
	}
	tmp := make([]*yunweiclient.GetClusterByPlatformData, 0)
	err = copier.Copy(&tmp, ens)
	if err != nil {
		return nil, xerr.NewErrMsg("复制平台集群信息失败，原因：" + err.Error())
	}

	return &yunweiclient.GetClusterByPlatformResp{
		Data: tmp,
	}, nil
}
