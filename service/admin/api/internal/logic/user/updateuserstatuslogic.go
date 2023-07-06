package user

import (
	"context"
	"encoding/json"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserStatusLogic) UpdateUserStatus(req *types.UserStatusReq) (err error) {
	_, err = l.svcCtx.AdminRpc.UpdateUserStatus(l.ctx, &admin.UserStatusReq{
		Id:           req.Id,
		Status:       req.Status,
		LastUpdateBy: ctxdata.GetUnameFromCtx(l.ctx),
	})
	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("修改用户状态失败,参数:%s,异常:%s", reqStr, err.Error())
		return err
	}

	return nil
}
