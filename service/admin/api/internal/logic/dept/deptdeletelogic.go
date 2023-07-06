package dept

import (
	"context"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDeleteLogic {
	return &DeptDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptDeleteLogic) DeptDelete(req *types.DeleteDeptReq) error {
	_, err := l.svcCtx.AdminRpc.DeptDelete(l.ctx, &admin.DeptDeleteReq{
		Id: req.DeptId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据deptId: %d,删除部门异常:%s", req.DeptId, err.Error())
		return err
	}
	return nil
}
