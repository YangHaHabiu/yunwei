package project

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/rpc/admin"

	"github.com/jinzhu/copier"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectListLogic {
	return &ProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectListLogic) ProjectList(req *types.ListProjectReq) (*types.ListProjectResp, error) {
	resp, err := l.svcCtx.AdminRpc.ProjectList(l.ctx, &admin.ProjectListReq{
		Current:     req.Current,
		PageSize:    req.PageSize,
		ProjectEn:   req.ProjectEn,
		ProjectCn:   req.ProjectCn,
		Status:      req.Status,
		ProjectType: req.ProjectType,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询列表异常:%s", string(data), err.Error())
		return nil, err
	}

	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label: "项目名",
			Value: "projectCn",
			Types: "input",
		},
		{
			Label: "项目英文",
			Value: "projectEn",
			Types: "input",
		},
		{
			Label: "状态",
			Value: "status",
			Types: "select",
			Children: []*types.FilterList{
				{
					Label: "在线",
					Value: "-1",
				},
				{
					Label: "下线",
					Value: "1",
				},
			},
		},
	}

	list := make([]*types.ListProjectData, 0)
	copier.Copy(&list, resp.List)

	return &types.ListProjectResp{
		Rows:   list,
		Total:  resp.Total,
		Filter: filterList,
	}, nil
}
