package project

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectUpdateLogic {
	return &ProjectUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectUpdateLogic) ProjectUpdate(req *types.UpdateProjectReq) (err error) {
	_, err = l.svcCtx.AdminRpc.ProjectUpdate(l.ctx, &admin.ProjectUpdateReq{
		ProjectId:   req.ProjectId,
		GroupType:   req.GroupType,
		ProjectCn:   req.ProjectCn,
		ProjectEn:   req.ProjectEn,
		ProjectTeam: req.ProjectTeam,
		ProjectType: req.ProjectType,
		GroupDevQq:  req.GroupDevQq,
		GroupQq:     req.GroupQq,
		CompanyId:   req.CompanyId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return
}
