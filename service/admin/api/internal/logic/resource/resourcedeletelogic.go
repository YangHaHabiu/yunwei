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

type ResourceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceDeleteLogic {
	return &ResourceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceDeleteLogic) ResourceDelete(req *types.DeleteResourceReq) error {
	var tmp []*adminclient.CommonResourceData
	err := copier.Copy(&tmp, req.ResourceData)
	if err != nil {
		return xerr.NewErrMsg("拷贝删除数据失败，原因：" + err.Error())
	}
	_, err = l.svcCtx.AdminRpc.ResourceDelete(l.ctx, &adminclient.DeleteResourceReq{
		ResourceData: tmp,
	})

	if err != nil {
		return err
	}
	return nil
}
