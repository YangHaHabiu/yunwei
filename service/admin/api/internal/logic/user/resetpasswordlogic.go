package user

import (
	"context"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReSetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReSetPasswordLogic) ReSetPassword(req *types.ReSetPasswordReq) (err error) {
	_, err = l.svcCtx.AdminRpc.ReSetPassword(l.ctx, &admin.ReSetPasswordReq{
		Id:           req.Id,
		NewPassword:  req.NewPassword,
		LastUpdateBy: ctxdata.GetUnameFromCtx(l.ctx),
	})
	if err != nil {
		return err
	}

	return nil
}
