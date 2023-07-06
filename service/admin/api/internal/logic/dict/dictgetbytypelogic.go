package dict

import (
	"context"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictGetByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictGetByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictGetByTypeLogic {
	return &DictGetByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictGetByTypeLogic) DictGetByType(req *types.DictGetByTypeReq) (*types.ListDictResp, error) {
	resp, err := l.svcCtx.AdminRpc.DictList(l.ctx, &admin.DictListReq{
		Current:  1,
		PageSize: 15,
		Types:    req.DictType,
		Pid:      -2,
	})

	if err != nil {
		return nil, err
	}

	var list []*types.ListDictData

	copier.Copy(&list, resp.List)

	return &types.ListDictResp{
		Rows:  list,
		Total: resp.Total,
	}, nil

}
