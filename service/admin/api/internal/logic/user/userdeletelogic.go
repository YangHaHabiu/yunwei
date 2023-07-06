package user

import (
	"context"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.DeleteUserReq) (err error) {
	_, err = l.svcCtx.AdminRpc.UserDelete(l.ctx, &admin.UserDeleteReq{
		Id:           req.Id,
		LastUpdateBy: ctxdata.GetUnameFromCtx(l.ctx),
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据userId: %d,删除用户异常:%s", req.Id, err.Error())
		return err
	}
	return nil
}
