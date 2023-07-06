package logic

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResourceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceAddLogic {
	return &ResourceAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// resource rpc start
func (l *ResourceAddLogic) ResourceAdd(in *adminclient.AddResourceReq) (*adminclient.AddResourceResp, error) {

	err := l.svcCtx.LabelGlobalModel.TransactInsert(l.ctx, in)
	if err != nil {
		return nil, xerr.NewErrMsg("新增资源信息失败" + err.Error())
	}

	return &adminclient.AddResourceResp{}, nil
}
