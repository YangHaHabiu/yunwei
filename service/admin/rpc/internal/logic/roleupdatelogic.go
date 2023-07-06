package logic

import (
	"context"
	"ywadmin-v3/common/ctxdata"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *adminclient.RoleUpdateReq) (*adminclient.RoleUpdateResp, error) {
	err := l.svcCtx.RoleModel.Update(l.ctx, &model.SysRole{
		Id:           in.Id,
		Name:         in.Name,
		Remark:       in.Remark,
		CreateBy:     ctxdata.GetUnameFromCtx(l.ctx),
		LastUpdateBy: in.LastUpdateBy,
		DelFlag:      0,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return &adminclient.RoleUpdateResp{}, nil
}
