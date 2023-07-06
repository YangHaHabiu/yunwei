package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LabelListByPriLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLabelListByPriLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LabelListByPriLogic {
	return &LabelListByPriLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LabelListByPriLogic) LabelListByPri(in *adminclient.LabelListByPriReq) (*adminclient.LabelListByPriResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "view_project_id__in", in.ProjectIds)

	var (
		err error
		all *[]model.LabelViewNew
	)

	all, err = l.svcCtx.LabelModel.FindAllClusterByPri(l.ctx, filters...)
	if err != nil {
		return nil, xerr.NewErrMsg("查询集群信息失败，原因：" + err.Error())
	}

	var list []*adminclient.LabelListByPriData
	err = copier.Copy(&list, all)
	if err != nil {
		return nil, xerr.NewErrMsg("复制集群信息失败，原因：" + err.Error())
	}
	return &adminclient.LabelListByPriResp{
		List: list,
	}, nil
}
