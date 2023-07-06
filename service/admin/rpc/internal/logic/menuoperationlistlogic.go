package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuOperationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuOperationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuOperationListLogic {
	return &MenuOperationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuOperationListLogic) MenuOperationList(in *adminclient.MenuOperationListReq) (*adminclient.MenuOperationListResp, error) {

	list, err := l.svcCtx.MenuModel.FindAllOperationList(l.ctx, &adminclient.MenuOperationListReq{})
	if err != nil {
		return nil, xerr.NewErrMsg("查询任务操作失败，原因：" + err.Error())
	}

	tmp := make([]*adminclient.MenuOperationListData, 0)

	err = copier.Copy(&tmp, list)
	if err != nil {
		return nil, xerr.NewErrMsg("复制任务操作信息失败，原因：" + err.Error())
	}

	return &adminclient.MenuOperationListResp{
		MenuOperationListData: tmp,
	}, nil
}
