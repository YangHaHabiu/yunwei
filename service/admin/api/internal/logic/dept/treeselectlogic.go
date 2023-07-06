package dept

import (
	"context"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
	"ywadmin-v3/service/admin/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeselectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTreeselectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeselectLogic {
	return &TreeselectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TreeselectLogic) Treeselect() (*types.TreeselectResp, error) {

	resp, err := l.svcCtx.AdminRpc.DeptList(l.ctx, &admin.DeptListReq{})
	if err != nil {
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

	return &types.TreeselectResp{
		TreeData: list,
	}, nil
}
