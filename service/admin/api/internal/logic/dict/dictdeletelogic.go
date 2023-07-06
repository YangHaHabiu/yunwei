package dict

import (
	"context"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDeleteLogic {
	return &DictDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictDeleteLogic) DictDelete(req *types.DeleteDictReq) error {
	list, err := l.svcCtx.AdminRpc.DictList(l.ctx, &adminclient.DictListReq{
		Current:  0,
		PageSize: 0,
		Pid:      req.DictId,
	})
	if err != nil {
		return err
	}
	if len(list.List) > 0 {
		return xerr.NewErrMsg("存在字典配置值")
	}
	_, err = l.svcCtx.AdminRpc.DictDelete(l.ctx, &admin.DictDeleteReq{
		Id: req.DictId,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据dictId: %d,删除字典异常:%s", req.DictId, err.Error())
		return err
	}

	return nil
}
