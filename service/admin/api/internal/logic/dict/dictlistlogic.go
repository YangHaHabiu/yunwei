package dict

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictListLogic {
	return &DictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictListLogic) DictList(req *types.ListDictReq) (*types.ListDictResp, error) {
	resp, err := l.svcCtx.AdminRpc.DictList(l.ctx, &admin.DictListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Types:    req.Types,
		Pid:      req.Pid,
		Id:       req.Id,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询字典列表异常:%s", string(data), err.Error())
		return nil, err
	}

	list := make([]*types.ListDictData, 0)

	copier.Copy(&list, resp.List)

	return &types.ListDictResp{
		Rows:  list,
		Total: resp.Total,
	}, nil
}
