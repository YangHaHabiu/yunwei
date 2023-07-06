package project

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectAddLogic {
	return &ProjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectAddLogic) ProjectAdd(req *types.AddProjectReq) (err error) {
	_, err = l.svcCtx.AdminRpc.ProjectAdd(l.ctx, &adminclient.ProjectAddReq{
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
		logx.WithContext(l.ctx).Errorf("添加信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}
	return
}
