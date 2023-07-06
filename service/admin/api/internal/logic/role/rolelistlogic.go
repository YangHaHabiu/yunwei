package role

import (
	"context"
	"encoding/json"
	"strconv"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.ListRoleReq) (*types.ListRoleResp, error) {
	resp, err := l.svcCtx.AdminRpc.RoleList(l.ctx, &admin.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询角色列表异常:%s", string(data), err.Error())
		return nil, err
	}

	list := make([]*types.ListRoleData, 0)

	for _, role := range resp.List {
		list = append(list, &types.ListRoleData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime,
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime,
			Label:          role.Name,
			Value:          strconv.FormatInt(role.Id, 10),
		})
	}

	return &types.ListRoleResp{
		Rows:  list,
		Total: resp.Total,
	}, nil
}
