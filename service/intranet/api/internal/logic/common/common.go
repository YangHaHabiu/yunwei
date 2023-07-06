package common

import (
	"context"
	"fmt"
	"strings"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/intranet/api/internal/svc"
	"ywadmin-v3/service/intranet/api/internal/types"

	"github.com/gogf/gf/util/gconv"
)

//获取个人项目值及列表
func GetProjectStrAndList(svcCtx *svc.ServiceContext, ctx context.Context, ProjectIds string) (string, []*types.FilterList, error) {
	ownerList, err := svcCtx.AdminRpc.ProjectOwnerList(ctx, &adminclient.ProjectOwnerReq{UserId: ctxdata.GetUidFromCtx(ctx)})
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

func GetinsideTasksStatus(svcCtx *svc.ServiceContext, ctx context.Context) ([]*types.FilterList, error) {
	return GetDictListByTypesNews(svcCtx, ctx, "inside_tasks_status", "内网任务状态")
}

//根据字典类型获取字典信息
func GetDictListByTypes(svcCtx *svc.ServiceContext, ctx context.Context, dictTypes, dictLabel string) ([]*types.FilterList, error) {
	TmpList := make([]*types.FilterList, 0)
	dictList, err := svcCtx.AdminRpc.DictList(ctx, &adminclient.DictListReq{
		Types:    dictTypes,
		Current:  0,
		PageSize: 0,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("获取" + dictLabel + "列表失败，原因：" + err.Error())
	}
	for _, v := range dictList.List {
		if v.Label == dictLabel {
			TmpList = append(TmpList, &types.FilterList{
				Label: v.Label,
				Value: v.Value,
			})
		}
	}

	return TmpList, nil
}

//根据字典类型获取字典信息
func GetDictListByTypesNews(svcCtx *svc.ServiceContext, ctx context.Context, dictTypes, dictLabel string) ([]*types.FilterList, error) {
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
