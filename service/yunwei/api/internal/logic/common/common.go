package common

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/yunwei/api/internal/svc"
	"ywadmin-v3/service/yunwei/api/internal/types"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/gogf/gf/util/gconv"
)

//获取个人项目值及列表
func GetProjectStrAndList(svcCtx *svc.ServiceContext, ctx context.Context, ProjectIds string, projectType ...string) (string, []*types.FilterList, error) {
	var typex string
	if len(projectType) > 0 {
		typex = projectType[0]
	}
	ownerList, err := svcCtx.AdminRpc.ProjectOwnerList(ctx, &adminclient.ProjectOwnerReq{UserId: ctxdata.GetUidFromCtx(ctx), ProjectType: typex})
	if err != nil {
		return "", nil, xerr.NewErrMsg("获取个人项目值及列表失败，原因：" + err.Error())
	}
	var projectIds string
	projectList := make([]*types.FilterList, 0)
	ownerTemps := make([]string, 0)
	ownerTemps = append(ownerTemps, "-1")
	for _, v := range ownerList.List {
		projectList = append(projectList, &types.FilterList{
			Label: v.ProjectCn,
			Value: gconv.String(v.ProjectId),
		})
		ownerTemps = append(ownerTemps, gconv.String(v.ProjectId))
	}
	if ProjectIds == "" {
		projectIds = ctxdata.GetProjectIdsFromCtx(ctx)
		if projectIds == "" {
			//根据用户id查询对应的项目
			projectIds = strings.Join(ownerTemps, ",")
		}
	} else {
		projectIds = ProjectIds
	}
	return projectIds, projectList, nil
}

//获取平台列表
func GetPlatform(svcCtx *svc.ServiceContext, ctx context.Context, projectIds, crossTypes, platformType string) ([]*types.FilterList, error) {
	platformList := make([]*types.FilterList, 0)
	if projectIds != "" {
		Platforms, err := svcCtx.YunWeiRpc.PlatformList(ctx, &yunweiclient.ListPlatformReq{
			Current:      0,
			PageSize:     0,
			DelFlag:      1,
			ProjectIds:   projectIds,
			Types:        crossTypes,
			PlatformType: platformType,
		})
		if err != nil {
			return nil, xerr.NewErrMsg("获取平台列表失败，原因：" + err.Error())
		}
		for _, v := range Platforms.Rows {
			platformList = append(platformList, &types.FilterList{
				Label: fmt.Sprintf("%s(%s)", v.PlatformEn, v.PlatformCn),
				Value: fmt.Sprintf("%d_%d", v.ProjectId, v.PlatformId),
			})
		}

	}
	return platformList, nil

}

//获取集群列表
func GetCluster(svcCtx *svc.ServiceContext, ctx context.Context, projectIds string) ([]*types.FilterList, error) {
	clusterList := make([]*types.FilterList, 0)
	clusters, err := svcCtx.AdminRpc.LabelListByPri(ctx, &adminclient.LabelListByPriReq{
		ProjectIds: projectIds,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("获取集群信息失败，原因：" + err.Error())
	}

	for _, v := range clusters.List {
		clusterList = append(clusterList, &types.FilterList{
			Label: v.ViewLabelName,
			Value: gconv.String(v.ViewLabelId),
		})
	}

	return clusterList, nil

}

//获取用户列表
func GetUserList(svcCtx *svc.ServiceContext, ctx context.Context, projectIds string) ([]*types.FilterList, error) {
	users, err := svcCtx.AdminRpc.UserList(ctx, &adminclient.UserListReq{
		Current: 0, PageSize: 0, ProjectIds: projectIds,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户列表失败，原因：" + err.Error())
	}
	userList := make([]*types.FilterList, 0)
	for _, v := range users.List {
		userList = append(userList, &types.FilterList{
			Label: fmt.Sprintf("%s(%s)", v.NickName, v.Name),
			Value: gconv.String(v.Id),
		})
	}
	return userList, nil

}

//获取出机方列表
func GetCompanyList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	companies, err := svcCtx.AdminRpc.CompanyList(ctx, &adminclient.CompanyListReq{
		Current: 0, PageSize: 0, SupplyCompanyStatus: "1",
	})
	if err != nil {
		return nil, xerr.NewErrMsg("获取出机方列表失败，原因：" + err.Error())
	}
	companyList := make([]*types.FilterList, 0)
	for _, v := range companies.List {
		companyList = append(companyList, &types.FilterList{
			Label: v.CompanyCn,
			Value: gconv.String(v.CompanyId),
		})
	}
	return companyList, nil
}

//获取云商列表
func GetProviderList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "cloud_provider_type", "云商")
}

//获取用途列表
func GetHostRoleList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "host_role_type", "用途")
}

//获取功能服列表
func GetFeatureServerList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "feature_server_type", "功能服")
}

//获取开服计划列表安装状态
func GetInstallStatusList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "install_status", "开服计划安装状态")
}

//获取开服计划列表清挡状态
func GetInitdbStatusList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "initdb_status", "开服计划清挡状态")
}

//获取合服计划列表安装状态
func GetMergeStatusList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "merge_status", "合服计划安装状态")
}

//获取计划队列列表状态
func GetScheduleTypeList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "schedule_types", "计划队列类型")
}

//获取计划队列列表状态
func GetScheduleStatusList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "schedule_status", "计划队列状态")
}

//获取阈值管理类型
func GetThresholdManageTypesList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypes(svcCtx, ctx, "threshold_manage_types", "阈值管理类型")
}

//根据字典类型获取字典信息
func GetDictListByTypes(svcCtx *svc.ServiceContext, ctx context.Context, dictTypes, dictLabel string) ([]*types.FilterList, error) {
	TmpList := make([]*types.FilterList, 0)
	dictList, err := svcCtx.AdminRpc.DictList(ctx, &adminclient.DictListReq{
		Pid:      -2,
		Types:    dictTypes,
		Current:  0,
		PageSize: 0,
	})

	if err != nil {
		return nil, xerr.NewErrMsg("获取" + dictLabel + "列表失败，原因：" + err.Error())
	}
	for _, v := range dictList.List {
		if dictTypes == "feature_server_type" {
			TmpList = append(TmpList, &types.FilterList{
				Label: v.Description,
				Value: v.Label,
			})
		} else if dictTypes == "host_role_type" {
			TmpList = append(TmpList, &types.FilterList{
				Label: v.Description,
				Value: v.Description,
			})
		} else {
			var tips string
			if v.Label != v.Description {
				tips = fmt.Sprintf("%s(%s)", v.Description, v.Label)
			} else {
				tips = v.Label
			}
			TmpList = append(TmpList, &types.FilterList{
				Label: tips,
				Value: v.Value,
			})
		}

	}
	return TmpList, nil
}

//获取计划字典类型
func GetScheduleType(svcCtx *svc.ServiceContext, ctx context.Context) (error, string) {
	//获取计划类型字典
	maps := make(map[string]string, 0)
	dictList, err := svcCtx.AdminRpc.DictList(ctx, &adminclient.DictListReq{
		Pid:      -2,
		Types:    "schedule_types",
		Current:  0,
		PageSize: 0,
	})
	if err != nil {
		return err, ""
	}
	//根据不同的键值确定对应关系表
	for _, v := range dictList.List {
		maps[v.Value] = v.Description
	}
	//序列化字典
	b, err := json.Marshal(maps)
	if err != nil {
		return err, ""

	}
	return nil, string(b)

}
