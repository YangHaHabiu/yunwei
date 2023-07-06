package user

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.ListUserReq) (*types.ListUserResp, error) {

	resp, err := l.svcCtx.AdminRpc.UserList(l.ctx, &adminclient.UserListReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		Status:     gconv.Int64(req.Status),
		DeptIds:    req.DeptIds,
		Email:      req.Email,
		Mobile:     req.Mobile,
		Name:       req.Name,
		NickName:   req.NickName,
		UserId:     req.Id,
		ProjectIds: req.ProjectIds,
		RoleIds:    req.RoleIds,
		UgroupIds:  req.UgroupIds,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询用户列表异常:%s", string(data), err.Error())
		return nil, err
	}

	var list []*types.ListUserData

	for _, item := range resp.List {
		listUserData := types.ListUserData{}
		_ = copier.Copy(&listUserData, &item)
		list = append(list, &listUserData)
	}

	tmpDeptList := make([]*types.FilterList, 0)
	deptList, err := l.svcCtx.AdminRpc.DeptList(l.ctx, &adminclient.DeptListReq{})
	if err != nil {
		return nil, err
	}
	for _, v := range deptList.List {
		tmpDeptList = append(tmpDeptList, &types.FilterList{
			Label: v.Name,
			Value: gconv.String(v.Id),
		})

	}
	tmpProjectList := make([]*types.FilterList, 0)
	projectList, err := l.svcCtx.AdminRpc.ProjectList(l.ctx, &adminclient.ProjectListReq{
		Current: 0, PageSize: 0, Status: "-1",
	})
	if err != nil {
		return nil, err
	}
	for _, v := range projectList.List {
		tmpProjectList = append(tmpProjectList, &types.FilterList{
			Label: v.ViewProjectCn + "(" + v.ViewCompanyEn + ")",
			Value: gconv.String(v.ViewProjectId),
		})
	}

	tmpRoleList := make([]*types.FilterList, 0)
	roleList, err := l.svcCtx.AdminRpc.RoleList(l.ctx, &adminclient.RoleListReq{
		Current: 0, PageSize: 0,
	})
	if err != nil {
		return nil, err
	}
	for _, v := range roleList.List {
		tmpRoleList = append(tmpRoleList, &types.FilterList{
			Label: v.Remark + "(" + v.Name + ")",
			Value: gconv.String(v.Id),
		})
	}

	tmpUgroupList := make([]*types.FilterList, 0)
	ugroupList, err := l.svcCtx.AdminRpc.UgroupList(l.ctx, &adminclient.UgroupListReq{
		Current: 0, PageSize: 0,
	})
	if err != nil {
		return nil, err
	}
	for _, v := range ugroupList.List {
		tmpUgroupList = append(tmpUgroupList, &types.FilterList{
			Label: v.UgName,
			Value: gconv.String(v.Id),
		})
	}

	//自定义筛选条件
	filterList := []*types.FilterList{
		{
			Label: "帐号",
			Value: "name",
			Types: "input",
		},
		{
			Label: "姓名",
			Value: "nickName",
			Types: "input",
		},
		{
			Label:    "部门",
			Value:    "deptIds",
			Types:    "select",
			Children: tmpDeptList,
		},
		{
			Label:    "项目",
			Value:    "projectIds",
			Types:    "select",
			Children: tmpProjectList,
		},
		{
			Label:    "角色",
			Value:    "roleIds",
			Types:    "select",
			Children: tmpRoleList,
		},
		{
			Label:    "用户组",
			Value:    "ugroupIds",
			Types:    "select",
			Children: tmpUgroupList,
		},
		{
			Label: "用户状态",
			Value: "status",
			Types: "select",
			Children: []*types.FilterList{
				{
					Label: "停用",
					Value: "2",
				},
				{
					Label: "启用",
					Value: "1",
				},
			},
		},
	}

	return &types.ListUserResp{
		Rows:   list,
		Total:  resp.Total,
		Filter: filterList,
	}, nil
}
