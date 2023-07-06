package platform

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlatformListByProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlatformListByProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlatformListByProjectLogic {
	return &PlatformListByProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlatformListByProjectLogic) PlatformListByProject(req *types.PlatformListByProjectReq) (resp *types.PlatformListByProjectResp, err error) {
	var projectIds string
	if req.ProjectId == 0 {
		projectIds = "-1"
	} else {
		projectIds = gconv.String(req.ProjectId)
	}

	list, err := l.svcCtx.YunWeiRpc.PlatformList(l.ctx, &yunweiclient.ListPlatformReq{
		Current:    0,
		PageSize:   0,
		ProjectIds: projectIds,
		Types:      req.NotType,
		DelFlag:    1,
	})
	if err != nil {
		return nil, err
	}

	tmp := make([]*types.PlatformCommon, 0)
	err = copier.Copy(&tmp, list.Rows)
	if err != nil {
		return nil, xerr.NewErrMsg("复制项目查询平台信息失败，原因：" + err.Error())
	}
	resp = new(types.PlatformListByProjectResp)
	resp.PlatfromData = tmp

	return
}
