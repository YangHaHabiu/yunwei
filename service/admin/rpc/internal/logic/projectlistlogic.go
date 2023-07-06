package logic

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectListLogic {
	return &ProjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProjectListLogic) ProjectList(in *adminclient.ProjectListReq) (*adminclient.NewProjectListResp, error) {

	var (
		count int64
		list  []*adminclient.NewProjectListData
		err   error
		all   *[]model.ViewCompanyProject
	)
	filters := make([]interface{}, 0)

	filters = append(filters,
		"view_project_id__in", in.ProjectIds,
		"view_project_cn__like", in.ProjectCn,
		"view_project_en__like", in.ProjectEn,
	)

	if in.ProjectType != "all" {
		filters = append(filters,
			"view_project_del_flag__=", in.Status,
		)
	}

	if in.Current != 0 && in.PageSize != 0 {
		count, _ = l.svcCtx.ProjectModel.Count(l.ctx, filters...)
		all, err = l.svcCtx.ProjectModel.FindPageListByPage(l.ctx, in.Current, in.PageSize, filters...)
	} else {
		all, err = l.svcCtx.ProjectModel.FindAll(l.ctx, filters...)
	}
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, xerr.NewErrMsg("查询列表信息失败" + err.Error())
	}

	for _, data := range *all {
		list = append(list, &adminclient.NewProjectListData{
			ViewCompanyId:      data.ViewCompanyId.Int64,
			ViewCompanyCn:      data.ViewCompanyCn.String,
			ViewCompanyEn:      data.ViewCompanyEn.String,
			ViewCompanyDelFlag: data.ViewCompanyDelFlag.Int64,
			ViewPrId:           data.ViewPrId.Int64,
			ViewProjectId:      data.ViewProjectId.Int64,
			ViewProjectCn:      data.ViewProjectCn.String,
			ViewProjectEn:      data.ViewProjectEn.String,
			ViewDeptId:         data.ViewDeptId.Int64,
			ViewDeptName:       data.ViewDeptName.String,
			ViewProjectType:    data.ViewProjectType.String,
			ViewGroupQq:        data.ViewGroupQq.String,
			ViewGroupTypeCn:    data.ViewGroupTypeCn.String,
			ViewGroupTypeEn:    data.ViewGroupTypeEn.String,
			ViewGroupDevQq:     data.ViewGroupDevQq.String,
			ViewProjectDelFlag: data.ViewProjectDelFlag.String,
			ViewProjectTypeCn:  data.ViewProjectTypeCn.String,
		})
	}
	return &adminclient.NewProjectListResp{
		Total: count,
		List:  list,
	}, nil
}
