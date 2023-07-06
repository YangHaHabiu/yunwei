package dept

import (
	"context"
	"encoding/json"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
	"ywadmin-v3/service/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptListLogic) DeptList(req *types.ListDeptReq) (*types.ListDeptResp, error) {
	resp, err := l.svcCtx.AdminRpc.DeptList(l.ctx, &admin.DeptListReq{
		Name:     req.Name,
		CreateBy: req.CreateBy,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询机构列表异常:%s", string(data), err.Error())
		return nil, err
	}

	list := make([]*types.ListDeptData, 0)

	for _, dept := range resp.List {
		list = append(list, &types.ListDeptData{
			Id:             dept.Id,
			Name:           dept.Name,
			ParentId:       dept.ParentId,
			OrderNum:       dept.OrderNum,
			CreateBy:       dept.CreateBy,
			CreateTime:     dept.CreateTime,
			LastUpdateBy:   dept.LastUpdateBy,
			LastUpdateTime: dept.LastUpdateTime,
		})
	}
	return &types.ListDeptResp{
		Rows:  list,
		Total: resp.Total,
	}, nil
}
