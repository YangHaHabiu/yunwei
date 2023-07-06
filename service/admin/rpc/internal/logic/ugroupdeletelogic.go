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

type UgroupDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUgroupDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UgroupDeleteLogic {
	return &UgroupDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UgroupDeleteLogic) UgroupDelete(in *adminclient.UgroupDeleteReq) (*adminclient.UgroupDeleteResp, error) {
	filters := make([]interface{}, 0)
	filters = append(filters, "ugroup_id__=", in.Id)
	all, err := l.svcCtx.UserUgroupModel.FindAll(l.ctx, filters...)
	if err != nil || len(*all) != 0 {
		return nil, xerr.NewErrMsg("用户有关联用户组数据，禁止删除，请检查")
	}
	var lastUpdateBy string
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		lastUpdateBy = md.Get("uname")[0]
	}

	err = l.svcCtx.UgroupModel.TransactDelete(l.ctx, &model.SysUgroup{Id: in.Id, LastUpdateBy: lastUpdateBy})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_DATA_DELETE_ERROR)
	}
	return &adminclient.UgroupDeleteResp{}, nil
}
