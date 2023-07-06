package logic

import (
	"context"
	"google.golang.org/grpc/metadata"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/model"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDeleteLogic) RoleDelete(in *adminclient.RoleDeleteReq) (*adminclient.RoleDeleteResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "role_id__=", in.Id)
	all, err := l.svcCtx.UserRoleModel.FindAll(l.ctx, filters...)
	if err != nil || len(*all) != 0 {
		return nil, xerr.NewErrMsg("角色关联用户数据，禁止删除，请检查")
	}
	var lastUpdateBy string
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		lastUpdateBy = md.Get("uname")[0]
	}

	err = l.svcCtx.RoleModel.TransactDelete(l.ctx, &model.SysRole{Id: in.Id, LastUpdateBy: lastUpdateBy})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_DELETE_ERROR)
	}
	return &adminclient.RoleDeleteResp{}, nil
}
