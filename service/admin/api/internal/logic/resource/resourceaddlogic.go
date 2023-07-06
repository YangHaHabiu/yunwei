package resource

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceAddLogic {
	return &ResourceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceAddLogic) ResourceAdd(req *types.AddResourceReq) error {
	var tmp []*adminclient.CommonResourceData
	err := copier.Copy(&tmp, req.ResourceData)
	if err != nil {
		return xerr.NewErrMsg("拷贝数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.AdminRpc.ResourceAdd(l.ctx, &adminclient.AddResourceReq{ResourceData: tmp})
	if err != nil {
		return err
	}
	return nil
}
