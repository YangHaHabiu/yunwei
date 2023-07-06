package logic

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/tool"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/yunwei/rpc/internal/svc"
	"ywadmin-v3/service/yunwei/rpc/yunweiclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigFileDeliveryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigFileDeliveryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigFileDeliveryListLogic {
	return &ConfigFileDeliveryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigFileDeliveryListLogic) ConfigFileDeliveryList(in *yunweiclient.ListConfigFileDeliveryReq) (*yunweiclient.ListConfigFileDeliveryResp, error) {
	list := make([]*yunweiclient.ListConfigFileData, 0)
	listNew := make([]*yunweiclient.ListConfigFileDeliveryDataTree, 0)
	all, err := l.svcCtx.ConfigFileModel.FindAll(l.ctx, "A.project_id__in", in.ProjectIds)
	if err != nil {
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}
	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("拷贝列表信息失败，原因：" + err.Error())
	}

	allList, err := l.svcCtx.ConfigMngLogModel.FindAllList(l.ctx, "view_user_project_id__in", in.ProjectIds)
	if err != nil {
		return nil, xerr.NewErrMsg("查询列表信息失败，原因：" + err.Error())
	}
	for _, v := range *allList {
		s := make([]string, 0)
		s, _ = tool.GetAllFile(l.svcCtx.Config.TemplateFilePath, s)
		listNew = append(listNew, &yunweiclient.ListConfigFileDeliveryDataTree{
			ProjectId: gconv.Int64(v.ViewUserProjectId),
			TotalList: v.AssetIps,
			MouldFile: s,
		})

	}
	//fmt.Println(list)

	//fmt.Println(listNew)
	return &yunweiclient.ListConfigFileDeliveryResp{
		Rows:      list,
		MergeRows: listNew,
	}, nil
}
