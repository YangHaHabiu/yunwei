package logic

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"strings"
	"ywadmin-v3/common/slicePage"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceObjectValueListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResourceObjectValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceObjectValueListLogic {
	return &ResourceObjectValueListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResourceObjectValueListLogic) ResourceObjectValueList(in *adminclient.ResourceObjectValueListReq) (*adminclient.ResourceObjectValueListResp, error) {

	var (
		primaryKeyValue     string
		bindPrimaryKeyValue string
		allResources        *[]model.ViewSearchLabel
		bindResources       *[]model.ViewSearchLabel
		listResources       *[]model.ViewSearchLabel
		all                 *[]model.LabelGlobalAndLabel
		err                 error
		alltmp              []*adminclient.ResourceObjectValueData
		bingTmp             []*adminclient.ResourceObjectValueData
		listTmp             []*adminclient.ResourceObjectValueData
		total               int64
	)

	if in.LabelType != "" && in.ResourceEnName != "" {
		all, err = l.svcCtx.LabelGlobalModel.FindAll(l.ctx,
			"l.label_id__=", in.LabelId,
			"lg.resource_en__=", in.ResourceEnName)
		if err != nil {
			return nil, xerr.NewErrMsg("查询资源ID信息失败，失败原因：" + err.Error())
		}
		*all = append(*all, model.LabelGlobalAndLabel{
			BindingId: -1,
		})
		bindPrimaryKeyValue = handleStr(all)
		if (in.LabelType == "1" && in.ResourceEnName == "platform") ||
			(in.LabelType == "1" && in.ResourceEnName == "asset") ||
			(in.LabelType == "3" && in.ResourceEnName == "platform") ||
			(in.LabelType == "2" && in.ResourceEnName == "platform") {
			//标签类型搜资源 （集群1，装服3）
			//1当前所有信息管标签类型
			//2当前标签类型已绑定的信息
			all, err = l.svcCtx.LabelGlobalModel.FindAll(l.ctx,
				"l.label_type__=", in.LabelType,
				"lg.resource_en__=", in.ResourceEnName)
			if err != nil {
				return nil, xerr.NewErrMsg("查询同类型资源信息失败，失败原因：" + err.Error())
			}
			*all = append(*all, model.LabelGlobalAndLabel{
				BindingId: -1,
			})

			primaryKeyValue = handleStr(all)
			if in.ResourceEnName == "asset" {
				allResources, err = l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx,
					"view_primary_key_value__not in", primaryKeyValue,
					"view_resource_en_name__=", in.ResourceEnName,
					"view_system_show__in", "1,3",
					"view_json_id->'$.view_recycle_type'__=", int64(2),
					"view_project_id__in", in.ProjectIds)

			} else {
				allResources, err = l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx,
					"view_primary_key_value__not in", primaryKeyValue,
					"view_resource_en_name__=", in.ResourceEnName,
					"view_system_show__in", "1,3",
					"view_project_id__in", in.ProjectIds)
			}

			bindResources, err = l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx,
				"view_primary_key_value__in", bindPrimaryKeyValue,
				"view_resource_en_name__=", in.ResourceEnName,
				"view_system_show__in", "1,3",
				"view_project_id__in", in.ProjectIds)
			*allResources = append(*allResources, *bindResources...)
		} else {

			//标签id搜资源 (功能服2，其他4)
			//1当前所有信息不管标签类型
			//2当前类型已绑定的信息
			//if in.ResourceEnName == "asset" {
			if in.LabelType == "3" && in.ResourceEnName == "asset" {
				allResources, err = l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx,
					"view_resource_en_name__=", in.ResourceEnName,
					"view_system_show__in", "1,3",
					"view_resource_type__in", "游戏服,跨服",
					"view_json_id->'$.view_recycle_type'__=", int64(2),
					"view_project_id__in", in.ProjectIds)
			} else {
				allResources, err = l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx,
					"view_resource_en_name__=", in.ResourceEnName,
					"view_system_show__in", "1,3",
					"view_json_id->'$.view_recycle_type'__=", int64(2),
					"view_project_id__in", in.ProjectIds)
			}

			//} else {
			//	allResources, err = l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx,
			//		"view_resource_en_name__=", in.ResourceEnName,
			//		"view_system_show__in", "1,3",
			//		"view_project_id__in", in.ProjectIds)
			//}
			bindResources, err = l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx,
				"view_primary_key_value__in", bindPrimaryKeyValue,
				"view_resource_en_name__=", in.ResourceEnName,
				"view_system_show__in", "1,3",
				"view_project_id__in", in.ProjectIds)

		}
	} else {
		allx, err := l.svcCtx.LabelGlobalModel.FindAllGroupByEnName(l.ctx, "l.label_id__=", in.LabelId)
		if err != nil {
			return nil, xerr.NewErrMsg("查询资源ID信息失败，失败原因：" + err.Error())
		}
		for _, v := range *allx {
			if v.LabelType != "2" {
				listResources, err = l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx,
					"view_resource_en_name__=", v.ResourceEn,
					"view_resource_en_name__like", in.ViewResourceEnName,
					"view_resource_value__regexp", in.ViewResourceValue,
					"view_json_id->'$.view_recycle_type'__=", gconv.Int64(in.ViewRecycleType),
					"view_primary_key_value__in", v.BindingIds,
					"view_system_show__in", "1,3",
					"view_project_id__in", in.ProjectIds)
			} else {
				listResources, err = l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx,
					"view_resource_en_name__=", v.ResourceEn,
					"view_resource_en_name__like", in.ViewResourceEnName,
					"view_resource_value__regexp", in.ViewResourceValue,
					"view_primary_key_value__in", v.BindingIds,
					"view_system_show__in", "1,3",
					"view_project_id__in", in.ProjectIds)
			}

			tmp := make([]*adminclient.ResourceObjectValueData, 0)
			err = copier.Copy(&tmp, listResources)
			if err != nil {
				return nil, xerr.NewErrMsg("拷贝所有类型所有资源失败，失败原因：" + err.Error())
			}
			listTmp = append(listTmp, tmp...)
		}
		total = int64(len(listTmp))
		if in.PageSize != 0 && in.Current != 0 {
			start, end := slicePage.SlicePage(in.Current, in.PageSize, total)
			listTmp = listTmp[start:end]
		}
	}

	if err != nil {
		return nil, xerr.NewErrMsg("查询资源类型失败，失败原因：" + err.Error())
	}
	if allResources != nil {
		err = copier.Copy(&alltmp, allResources)
		if err != nil {
			return nil, xerr.NewErrMsg("拷贝资源信息失败，失败原因1：" + err.Error())
		}
	}

	if bindResources != nil {
		err = copier.Copy(&bingTmp, bindResources)
		if err != nil {
			return nil, xerr.NewErrMsg("拷贝资源信息失败，失败原因2：" + err.Error())
		}
	}

	return &adminclient.ResourceObjectValueListResp{
		//单个类型所有资源
		List: alltmp,
		//单个类型绑定资源
		BingList: bingTmp,
		//绑定所有类型的列表绑定labelid
		AllList: listTmp,
		Total:   total,
	}, nil
}

func handleStr(all *[]model.LabelGlobalAndLabel) (keyValue string) {
	tmp := make([]string, 0)
	for _, v := range *all {
		tmp = append(tmp, gconv.String(v.BindingId))
	}

	if len(tmp) != 0 {
		return strings.Join(tmp, ",")
	}
	return ""

}
