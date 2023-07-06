package logic

import (
	"context"
	"ywadmin-v3/common/globalkey"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResourceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceListLogic {
	return &ResourceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResourceListLogic) ResourceList(in *adminclient.ListResourceReq) (*adminclient.ListResourceResp, error) {
	//all, err := l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx, "view_system_show__in", "1,3",
	//	"view_project_id__in", in.ProjectIds,
	//	globalkey.LabelTypes[in.LabelType]+"__=", "1")
	all, err := l.svcCtx.LabelGlobalModel.FindAllBindResource(l.ctx, "view_system_show__in", "1,3",
		globalkey.LabelTypes[in.LabelType]+"__=", "1")
	if err != nil {
		return nil, xerr.NewErrMsg("查询所属的资源组信息失败,原因：" + err.Error())
	}
	tmp := make([]*adminclient.ListResourceList, 0)
	tmpMap := make(map[string]interface{})
	for _, val := range *all {

		//判断主键为val的map是否存在
		if _, ok := tmpMap[val.ViewResourceEnName]; !ok {
			tmp = append(tmp, &adminclient.ListResourceList{
				Label: val.ViewResourceCnName,
				Value: val.ViewResourceEnName,
			})
			tmpMap[val.ViewResourceEnName] = nil
		}
	}

	return &adminclient.ListResourceResp{
		List: tmp,
	}, nil
}
