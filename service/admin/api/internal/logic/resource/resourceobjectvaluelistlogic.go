package resource

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"strings"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceObjectValueListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceObjectValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceObjectValueListLogic {
	return &ResourceObjectValueListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceObjectValueListLogic) ResourceObjectValueList(req *types.ResourceObjectValueListReq) (*types.ResourceObjectValueListResp, error) {
	ownerList, err := l.svcCtx.AdminRpc.ProjectOwnerList(l.ctx, &adminclient.ProjectOwnerReq{UserId: ctxdata.GetUidFromCtx(l.ctx)})
	if err != nil {
		return nil, err
	}

	ownerTmpx := make([]string, 0)
	ownerTmpx = append(ownerTmpx, "-1")
	for _, v := range ownerList.List {
		ownerTmpx = append(ownerTmpx, gconv.String(v.ProjectId))
	}
	list, err := l.svcCtx.AdminRpc.ResourceObjectValueList(l.ctx, &adminclient.ResourceObjectValueListReq{
		ResourceEnName:     req.ResourceEn,
		ProjectIds:         strings.Join(ownerTmpx, ","),
		LabelId:            req.LabelId,
		LabelType:          req.LabelType,
		Current:            req.Current,
		PageSize:           req.PageSize,
		ViewResourceValue:  req.ViewResourceValue,
		ViewResourceEnName: req.ViewResourceEnName,
		ViewRecycleType:    req.ViewRecycleType,
	})
	if err != nil {
		return nil, err
	}

	alllist := make([]*types.ResourceObjectValueListData, 0)
	bindlist := make([]*types.ResourceObjectValueListData, 0)
	lists := make([]*types.ResourceObjectValueListData, 0)

	err = copier.Copy(&alllist, list.AllList)
	if err != nil {
		return nil, xerr.NewErrMsg("复制所有类型绑定资源数据失败，原因：" + err.Error())
	}
	err = copier.Copy(&bindlist, list.BingList)
	if err != nil {
		return nil, xerr.NewErrMsg("复制单个资源绑定资源数据失败，原因：" + err.Error())
	}

	err = copier.Copy(&lists, list.List)
	if err != nil {
		return nil, xerr.NewErrMsg("复制单个类型所有资源数据失败，原因：" + err.Error())
	}

	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label: "表名",
			Value: "viewResourceEnName",
			Types: "input",
		},
		{
			Label: "数据值",
			Value: "viewResourceValue",
			Types: "input",
		},
	}
	if req.LabelType != "2" {
		filterList = append(filterList, &types.FilterList{
			Label: "回收状态",
			Value: "viewRecycleType",
			Types: "select",
			Children: []*types.FilterList{
				{
					Label: "回收",
					Value: "1",
				},
				{
					Label: "正常",
					Value: "2",
				},
			},
		})
	}

	return &types.ResourceObjectValueListResp{
		Rows:     lists,
		AllRows:  alllist,
		BindRows: bindlist,
		Total:    list.Total,
		Filter:   filterList}, nil
}
