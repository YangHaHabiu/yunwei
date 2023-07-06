package user

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/service/admin/api/internal/svc"
	"ywadmin-v3/service/admin/api/internal/types"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/admin/rpc/adminclient"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.GetUserByIdReq) (*types.GetUserByIdResp, error) {
	roleList, roleErr := l.svcCtx.AdminRpc.RoleList(l.ctx, &admin.RoleListReq{Current: 1, PageSize: 20})
	if roleErr != nil {
		return nil, xerr.NewErrMsg("获取用户角色信息失败")
	}
	var roleData []*types.RoleAllResp

	for _, role := range roleList.List {
		roleData = append(roleData, &types.RoleAllResp{
			Id:     role.Id,
			Name:   role.Name,
			Remark: role.Remark,
		})
	}

	uroupList, err := l.svcCtx.AdminRpc.UgroupList(l.ctx, &admin.UgroupListReq{Current: 1, PageSize: 20})
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户组信息失败")
	}

	var ugroupData []*types.UgroupAllResp
	for _, ugroup := range uroupList.List {
		ugroupData = append(ugroupData, &types.UgroupAllResp{
			Id:     ugroup.Id,
			Name:   ugroup.UgName,
			Remark: ugroup.UgJson,
		})
	}
	if req.UserId != 0 {

		list, err := l.svcCtx.AdminRpc.UserList(l.ctx, &adminclient.UserListReq{
			Current:  1,
			PageSize: 20,
			UserId:   req.UserId,
		})

		if err != nil {
			return nil, xerr.NewErrMsg("找不到该用户信息1")
		}
		if len(list.List) != 1 {
			return nil, xerr.NewErrMsg("找不到该用户信息2")
		}
		var userInfo types.ListUserData
		_ = copier.Copy(&userInfo, list.List[0])
		roleIds := make([]int64, 0)
		for _, v := range strings.Split(list.List[0].RoleIds, ",") {
			if gconv.Int64(v) != 0 {
				roleIds = append(roleIds, gconv.Int64(v))
			}
		}
		ugroupIds := make([]int64, 0)
		for _, v := range strings.Split(list.List[0].UgroupIds, ",") {
			if gconv.Int64(v) != 0 {
				ugroupIds = append(ugroupIds, gconv.Int64(v))
			}

		}
		return &types.GetUserByIdResp{
			Data:      userInfo,
			RoleIds:   roleIds,
			Roles:     roleData,
			UgroupIds: ugroupIds,
			Ugroups:   ugroupData,
		}, nil
	}

	return &types.GetUserByIdResp{
		Roles:   roleData,
		Ugroups: ugroupData,
	}, nil

}
