package user

import (
	"context"
	"strconv"
	"ywadmin-v3/service/admin/rpc/admin"

	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAllDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectAllDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectAllDataLogic {
	return &SelectAllDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectAllDataLogic) SelectAllData(req *types.SelectDataReq) (*types.SelectDataResp, error) {
	roleList, roleErr := l.svcCtx.AdminRpc.RoleList(l.ctx, &admin.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if roleErr != nil {
		return nil, roleErr
	}

	var roleData []*types.RoleAllResp

	for _, role := range roleList.List {
		roleData = append(roleData, &types.RoleAllResp{
			Id:     role.Id,
			Name:   role.Name,
			Remark: role.Remark,
		})
	}

	deptList, err := l.svcCtx.AdminRpc.DeptList(l.ctx, &admin.DeptListReq{})

	if err != nil {
		return nil, err
	}

	var deptData []*types.DeptAllResp

	for _, dept := range deptList.List {
		deptData = append(deptData, &types.DeptAllResp{
			Id:       dept.Id,
			Value:    strconv.FormatInt(dept.Id, 10),
			Title:    dept.Name,
			ParentId: dept.ParentId,
		})
	}

	uroupList, err := l.svcCtx.AdminRpc.UgroupList(l.ctx, &admin.UgroupListReq{Current: 1, PageSize: 20})
	if err != nil {
		return nil, err
	}

	var ugroupData []*types.UgroupAllResp
	for _, ugroup := range uroupList.List {
		ugroupData = append(ugroupData, &types.UgroupAllResp{
			Id:     ugroup.Id,
			Name:   ugroup.UgName,
			Remark: ugroup.UgJson,
		})
	}

	return &types.SelectDataResp{
		RoleAll:   roleData,
		DeptAll:   deptData,
		UgroupAll: ugroupData,
	}, nil
}
