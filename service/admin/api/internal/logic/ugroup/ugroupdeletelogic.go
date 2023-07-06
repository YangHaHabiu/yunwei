package ugroup

import (
	"context"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UgroupDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUgroupDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupDeleteLogic {
	return &UgroupDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UgroupDeleteLogic) UgroupDelete(req *types.DeleteUgroupReq) (err error) {
	_, err = l.svcCtx.AdminRpc.UgroupDelete(l.ctx, &admin.UgroupDeleteReq{
		Id: req.UgroupId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除异常:%s", req.UgroupId, err.Error())
		return err
	}
	return nil
}
