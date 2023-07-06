// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package server

import (
	"context"

	"ywadmin-v3/service/admin/rpc/adminclient"
	"ywadmin-v3/service/admin/rpc/internal/logic"
	"ywadmin-v3/service/admin/rpc/internal/svc"
)

type AdminServer struct {
	svcCtx *svc.ServiceContext
	adminclient.UnimplementedAdminServer
}

func NewAdminServer(svcCtx *svc.ServiceContext) *AdminServer {
	return &AdminServer{
		svcCtx: svcCtx,
	}
}

// user rpc start
func (s *AdminServer) Login(ctx context.Context, in *adminclient.LoginReq) (*adminclient.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *AdminServer) Logout(ctx context.Context, in *adminclient.LogoutReq) (*adminclient.LogoutResp, error) {
	l := logic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}

func (s *AdminServer) UserInfo(ctx context.Context, in *adminclient.InfoReq) (*adminclient.InfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *AdminServer) UserAdd(ctx context.Context, in *adminclient.UserAddReq) (*adminclient.UserAddResp, error) {
	l := logic.NewUserAddLogic(ctx, s.svcCtx)
	return l.UserAdd(in)
}

func (s *AdminServer) UserList(ctx context.Context, in *adminclient.UserListReq) (*adminclient.UserListResp, error) {
	l := logic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}

func (s *AdminServer) UserUpdate(ctx context.Context, in *adminclient.UserUpdateReq) (*adminclient.UserUpdateResp, error) {
	l := logic.NewUserUpdateLogic(ctx, s.svcCtx)
	return l.UserUpdate(in)
}

func (s *AdminServer) UserDelete(ctx context.Context, in *adminclient.UserDeleteReq) (*adminclient.UserDeleteResp, error) {
	l := logic.NewUserDeleteLogic(ctx, s.svcCtx)
	return l.UserDelete(in)
}

func (s *AdminServer) ReSetPassword(ctx context.Context, in *adminclient.ReSetPasswordReq) (*adminclient.ReSetPasswordResp, error) {
	l := logic.NewReSetPasswordLogic(ctx, s.svcCtx)
	return l.ReSetPassword(in)
}

func (s *AdminServer) UpdateUserStatus(ctx context.Context, in *adminclient.UserStatusReq) (*adminclient.UserStatusResp, error) {
	l := logic.NewUpdateUserStatusLogic(ctx, s.svcCtx)
	return l.UpdateUserStatus(in)
}

func (s *AdminServer) UpdatePersonalInfo(ctx context.Context, in *adminclient.UserUpdatePersonalInfoReq) (*adminclient.UserUpdatePersonalInfoResp, error) {
	l := logic.NewUpdatePersonalInfoLogic(ctx, s.svcCtx)
	return l.UpdatePersonalInfo(in)
}

func (s *AdminServer) UpdatePersonalPassword(ctx context.Context, in *adminclient.UserUpdatePersonalPasswordReq) (*adminclient.UserUpdatePersonalPasswordResp, error) {
	l := logic.NewUpdatePersonalPasswordLogic(ctx, s.svcCtx)
	return l.UpdatePersonalPassword(in)
}

func (s *AdminServer) UserStrategyList(ctx context.Context, in *adminclient.UserStrategyInfoReq) (*adminclient.UserStrategyInfoResp, error) {
	l := logic.NewUserStrategyListLogic(ctx, s.svcCtx)
	return l.UserStrategyList(in)
}

func (s *AdminServer) GetUserAssignmentPolicy(ctx context.Context, in *adminclient.GetUserAssignmentPolicyReq) (*adminclient.GetUserAssignmentPolicyResp, error) {
	l := logic.NewGetUserAssignmentPolicyLogic(ctx, s.svcCtx)
	return l.GetUserAssignmentPolicy(in)
}

func (s *AdminServer) UserAssignmentPolicy(ctx context.Context, in *adminclient.UserAssignmentPolicyReq) (*adminclient.UserAssignmentPolicyResp, error) {
	l := logic.NewUserAssignmentPolicyLogic(ctx, s.svcCtx)
	return l.UserAssignmentPolicy(in)
}

func (s *AdminServer) UserBatchEditItems(ctx context.Context, in *adminclient.UserBatchEditItemsReq) (*adminclient.UserBatchEditItemsResp, error) {
	l := logic.NewUserBatchEditItemsLogic(ctx, s.svcCtx)
	return l.UserBatchEditItems(in)
}

// ugroup rpc start
func (s *AdminServer) UgroupAdd(ctx context.Context, in *adminclient.UgroupAddReq) (*adminclient.UgroupAddResp, error) {
	l := logic.NewUgroupAddLogic(ctx, s.svcCtx)
	return l.UgroupAdd(in)
}

func (s *AdminServer) UgroupList(ctx context.Context, in *adminclient.UgroupListReq) (*adminclient.UgroupListResp, error) {
	l := logic.NewUgroupListLogic(ctx, s.svcCtx)
	return l.UgroupList(in)
}

func (s *AdminServer) UgroupUpdate(ctx context.Context, in *adminclient.UgroupUpdateReq) (*adminclient.UgroupUpdateResp, error) {
	l := logic.NewUgroupUpdateLogic(ctx, s.svcCtx)
	return l.UgroupUpdate(in)
}

func (s *AdminServer) UgroupDelete(ctx context.Context, in *adminclient.UgroupDeleteReq) (*adminclient.UgroupDeleteResp, error) {
	l := logic.NewUgroupDeleteLogic(ctx, s.svcCtx)
	return l.UgroupDelete(in)
}

func (s *AdminServer) UgroupInfo(ctx context.Context, in *adminclient.UgroupInfoReq) (*adminclient.UgroupInfoResp, error) {
	l := logic.NewUgroupInfoLogic(ctx, s.svcCtx)
	return l.UgroupInfo(in)
}

func (s *AdminServer) GetUgroupAssignmentPolicy(ctx context.Context, in *adminclient.GetUgroupAssignmentPolicyReq) (*adminclient.GetUgroupAssignmentPolicyResp, error) {
	l := logic.NewGetUgroupAssignmentPolicyLogic(ctx, s.svcCtx)
	return l.GetUgroupAssignmentPolicy(in)
}

func (s *AdminServer) UgroupAssignmentPolicy(ctx context.Context, in *adminclient.UgroupAssignmentPolicyReq) (*adminclient.UgroupAssignmentPolicyResp, error) {
	l := logic.NewUgroupAssignmentPolicyLogic(ctx, s.svcCtx)
	return l.UgroupAssignmentPolicy(in)
}

func (s *AdminServer) GetUgroupAssignmentUser(ctx context.Context, in *adminclient.GetUgroupAssignmentUserReq) (*adminclient.GetUgroupAssignmentUserResp, error) {
	l := logic.NewGetUgroupAssignmentUserLogic(ctx, s.svcCtx)
	return l.GetUgroupAssignmentUser(in)
}

func (s *AdminServer) UgroupAssignmentUser(ctx context.Context, in *adminclient.UgroupAssignmentUserReq) (*adminclient.UgroupAssignmentUserResp, error) {
	l := logic.NewUgroupAssignmentUserLogic(ctx, s.svcCtx)
	return l.UgroupAssignmentUser(in)
}

// stgroup rpc start
func (s *AdminServer) StgroupAdd(ctx context.Context, in *adminclient.StgroupAddReq) (*adminclient.StgroupAddResp, error) {
	l := logic.NewStgroupAddLogic(ctx, s.svcCtx)
	return l.StgroupAdd(in)
}

func (s *AdminServer) StgroupList(ctx context.Context, in *adminclient.StgroupListReq) (*adminclient.StgroupListResp, error) {
	l := logic.NewStgroupListLogic(ctx, s.svcCtx)
	return l.StgroupList(in)
}

func (s *AdminServer) StgroupUpdate(ctx context.Context, in *adminclient.StgroupUpdateReq) (*adminclient.StgroupUpdateResp, error) {
	l := logic.NewStgroupUpdateLogic(ctx, s.svcCtx)
	return l.StgroupUpdate(in)
}

func (s *AdminServer) StgroupDelete(ctx context.Context, in *adminclient.StgroupDeleteReq) (*adminclient.StgroupDeleteResp, error) {
	l := logic.NewStgroupDeleteLogic(ctx, s.svcCtx)
	return l.StgroupDelete(in)
}

func (s *AdminServer) StgroupInfo(ctx context.Context, in *adminclient.StgroupInfoReq) (*adminclient.StgroupInfoResp, error) {
	l := logic.NewStgroupInfoLogic(ctx, s.svcCtx)
	return l.StgroupInfo(in)
}

func (s *AdminServer) PolicyAssociatedUsers(ctx context.Context, in *adminclient.PolicyAssociatedUsersReq) (*adminclient.PolicyAssociatedUsersResp, error) {
	l := logic.NewPolicyAssociatedUsersLogic(ctx, s.svcCtx)
	return l.PolicyAssociatedUsers(in)
}

func (s *AdminServer) GetUserCheckStategyInfo(ctx context.Context, in *adminclient.StgroupUserCheckInfoReq) (*adminclient.StgroupUserCheckInfoResp, error) {
	l := logic.NewGetUserCheckStategyInfoLogic(ctx, s.svcCtx)
	return l.GetUserCheckStategyInfo(in)
}

// strategy rpc start
func (s *AdminServer) StrategyList(ctx context.Context, in *adminclient.StrategyListReq) (*adminclient.StrategyListResp, error) {
	l := logic.NewStrategyListLogic(ctx, s.svcCtx)
	return l.StrategyList(in)
}

// role rpc start
func (s *AdminServer) RoleAdd(ctx context.Context, in *adminclient.RoleAddReq) (*adminclient.RoleAddResp, error) {
	l := logic.NewRoleAddLogic(ctx, s.svcCtx)
	return l.RoleAdd(in)
}

func (s *AdminServer) RoleList(ctx context.Context, in *adminclient.RoleListReq) (*adminclient.RoleListResp, error) {
	l := logic.NewRoleListLogic(ctx, s.svcCtx)
	return l.RoleList(in)
}

func (s *AdminServer) RoleUpdate(ctx context.Context, in *adminclient.RoleUpdateReq) (*adminclient.RoleUpdateResp, error) {
	l := logic.NewRoleUpdateLogic(ctx, s.svcCtx)
	return l.RoleUpdate(in)
}

func (s *AdminServer) RoleDelete(ctx context.Context, in *adminclient.RoleDeleteReq) (*adminclient.RoleDeleteResp, error) {
	l := logic.NewRoleDeleteLogic(ctx, s.svcCtx)
	return l.RoleDelete(in)
}

func (s *AdminServer) QueryMenuByRoleId(ctx context.Context, in *adminclient.QueryMenuByRoleIdReq) (*adminclient.QueryMenuByRoleIdResp, error) {
	l := logic.NewQueryMenuByRoleIdLogic(ctx, s.svcCtx)
	return l.QueryMenuByRoleId(in)
}

func (s *AdminServer) UpdateMenuRole(ctx context.Context, in *adminclient.UpdateMenuRoleReq) (*adminclient.UpdateMenuRoleResp, error) {
	l := logic.NewUpdateMenuRoleLogic(ctx, s.svcCtx)
	return l.UpdateMenuRole(in)
}

func (s *AdminServer) GetRoleAssignmentUser(ctx context.Context, in *adminclient.GetRoleAssignmentUserReq) (*adminclient.GetRoleAssignmentUserResp, error) {
	l := logic.NewGetRoleAssignmentUserLogic(ctx, s.svcCtx)
	return l.GetRoleAssignmentUser(in)
}

func (s *AdminServer) RoleAssignmentUser(ctx context.Context, in *adminclient.RoleAssignmentUserReq) (*adminclient.RoleAssignmentUserResp, error) {
	l := logic.NewRoleAssignmentUserLogic(ctx, s.svcCtx)
	return l.RoleAssignmentUser(in)
}

// menu rpc start
func (s *AdminServer) MenuAdd(ctx context.Context, in *adminclient.MenuAddReq) (*adminclient.MenuAddResp, error) {
	l := logic.NewMenuAddLogic(ctx, s.svcCtx)
	return l.MenuAdd(in)
}

func (s *AdminServer) MenuList(ctx context.Context, in *adminclient.MenuListReq) (*adminclient.MenuListResp, error) {
	l := logic.NewMenuListLogic(ctx, s.svcCtx)
	return l.MenuList(in)
}

func (s *AdminServer) MenuUpdate(ctx context.Context, in *adminclient.MenuUpdateReq) (*adminclient.MenuUpdateResp, error) {
	l := logic.NewMenuUpdateLogic(ctx, s.svcCtx)
	return l.MenuUpdate(in)
}

func (s *AdminServer) MenuDelete(ctx context.Context, in *adminclient.MenuDeleteReq) (*adminclient.MenuDeleteResp, error) {
	l := logic.NewMenuDeleteLogic(ctx, s.svcCtx)
	return l.MenuDelete(in)
}

func (s *AdminServer) MenuOperationList(ctx context.Context, in *adminclient.MenuOperationListReq) (*adminclient.MenuOperationListResp, error) {
	l := logic.NewMenuOperationListLogic(ctx, s.svcCtx)
	return l.MenuOperationList(in)
}

// dict rpc start
func (s *AdminServer) DictAdd(ctx context.Context, in *adminclient.DictAddReq) (*adminclient.DictAddResp, error) {
	l := logic.NewDictAddLogic(ctx, s.svcCtx)
	return l.DictAdd(in)
}

func (s *AdminServer) DictList(ctx context.Context, in *adminclient.DictListReq) (*adminclient.DictListResp, error) {
	l := logic.NewDictListLogic(ctx, s.svcCtx)
	return l.DictList(in)
}

func (s *AdminServer) DictUpdate(ctx context.Context, in *adminclient.DictUpdateReq) (*adminclient.DictUpdateResp, error) {
	l := logic.NewDictUpdateLogic(ctx, s.svcCtx)
	return l.DictUpdate(in)
}

func (s *AdminServer) DictDelete(ctx context.Context, in *adminclient.DictDeleteReq) (*adminclient.DictDeleteResp, error) {
	l := logic.NewDictDeleteLogic(ctx, s.svcCtx)
	return l.DictDelete(in)
}

// dept rpc start
func (s *AdminServer) DeptAdd(ctx context.Context, in *adminclient.DeptAddReq) (*adminclient.DeptAddResp, error) {
	l := logic.NewDeptAddLogic(ctx, s.svcCtx)
	return l.DeptAdd(in)
}

func (s *AdminServer) DeptList(ctx context.Context, in *adminclient.DeptListReq) (*adminclient.DeptListResp, error) {
	l := logic.NewDeptListLogic(ctx, s.svcCtx)
	return l.DeptList(in)
}

func (s *AdminServer) DeptUpdate(ctx context.Context, in *adminclient.DeptUpdateReq) (*adminclient.DeptUpdateResp, error) {
	l := logic.NewDeptUpdateLogic(ctx, s.svcCtx)
	return l.DeptUpdate(in)
}

func (s *AdminServer) DeptDelete(ctx context.Context, in *adminclient.DeptDeleteReq) (*adminclient.DeptDeleteResp, error) {
	l := logic.NewDeptDeleteLogic(ctx, s.svcCtx)
	return l.DeptDelete(in)
}

// loginlog rpc start
func (s *AdminServer) LoginLogAdd(ctx context.Context, in *adminclient.LoginLogAddReq) (*adminclient.LoginLogAddResp, error) {
	l := logic.NewLoginLogAddLogic(ctx, s.svcCtx)
	return l.LoginLogAdd(in)
}

func (s *AdminServer) LoginLogUpdate(ctx context.Context, in *adminclient.LoginLogUpdateReq) (*adminclient.LoginLogUpdateResp, error) {
	l := logic.NewLoginLogUpdateLogic(ctx, s.svcCtx)
	return l.LoginLogUpdate(in)
}

func (s *AdminServer) LoginLogList(ctx context.Context, in *adminclient.LoginLogListReq) (*adminclient.LoginLogListResp, error) {
	l := logic.NewLoginLogListLogic(ctx, s.svcCtx)
	return l.LoginLogList(in)
}

// syslog rpc start
func (s *AdminServer) SysLogAdd(ctx context.Context, in *adminclient.SysLogAddReq) (*adminclient.SysLogAddResp, error) {
	l := logic.NewSysLogAddLogic(ctx, s.svcCtx)
	return l.SysLogAdd(in)
}

func (s *AdminServer) SysLogList(ctx context.Context, in *adminclient.SysLogListReq) (*adminclient.SysLogListResp, error) {
	l := logic.NewSysLogListLogic(ctx, s.svcCtx)
	return l.SysLogList(in)
}

// company rpc start
func (s *AdminServer) CompanyAdd(ctx context.Context, in *adminclient.CompanyAddReq) (*adminclient.CompanyAddResp, error) {
	l := logic.NewCompanyAddLogic(ctx, s.svcCtx)
	return l.CompanyAdd(in)
}

func (s *AdminServer) CompanyList(ctx context.Context, in *adminclient.CompanyListReq) (*adminclient.CompanyListResp, error) {
	l := logic.NewCompanyListLogic(ctx, s.svcCtx)
	return l.CompanyList(in)
}

func (s *AdminServer) CompanyUpdate(ctx context.Context, in *adminclient.CompanyUpdateReq) (*adminclient.CompanyUpdateResp, error) {
	l := logic.NewCompanyUpdateLogic(ctx, s.svcCtx)
	return l.CompanyUpdate(in)
}

func (s *AdminServer) CompanyDelete(ctx context.Context, in *adminclient.CompanyDeleteReq) (*adminclient.CompanyDeleteResp, error) {
	l := logic.NewCompanyDeleteLogic(ctx, s.svcCtx)
	return l.CompanyDelete(in)
}

func (s *AdminServer) UpdateSupplyCompany(ctx context.Context, in *adminclient.UpdateSupplyCompanyReq) (*adminclient.UpdateSupplyCompanyResp, error) {
	l := logic.NewUpdateSupplyCompanyLogic(ctx, s.svcCtx)
	return l.UpdateSupplyCompany(in)
}

// project rpc start
func (s *AdminServer) ProjectAdd(ctx context.Context, in *adminclient.ProjectAddReq) (*adminclient.ProjectAddResp, error) {
	l := logic.NewProjectAddLogic(ctx, s.svcCtx)
	return l.ProjectAdd(in)
}

func (s *AdminServer) ProjectList(ctx context.Context, in *adminclient.ProjectListReq) (*adminclient.NewProjectListResp, error) {
	l := logic.NewProjectListLogic(ctx, s.svcCtx)
	return l.ProjectList(in)
}

func (s *AdminServer) ProjectUpdate(ctx context.Context, in *adminclient.ProjectUpdateReq) (*adminclient.ProjectUpdateResp, error) {
	l := logic.NewProjectUpdateLogic(ctx, s.svcCtx)
	return l.ProjectUpdate(in)
}

func (s *AdminServer) ProjectDelete(ctx context.Context, in *adminclient.ProjectDeleteReq) (*adminclient.ProjectDeleteResp, error) {
	l := logic.NewProjectDeleteLogic(ctx, s.svcCtx)
	return l.ProjectDelete(in)
}

func (s *AdminServer) ProjectOwnerList(ctx context.Context, in *adminclient.ProjectOwnerReq) (*adminclient.ProjectListResp, error) {
	l := logic.NewProjectOwnerListLogic(ctx, s.svcCtx)
	return l.ProjectOwnerList(in)
}

func (s *AdminServer) ProjectGetOne(ctx context.Context, in *adminclient.ProjectGetOneReq) (*adminclient.ProjectListData, error) {
	l := logic.NewProjectGetOneLogic(ctx, s.svcCtx)
	return l.ProjectGetOne(in)
}

// label rpc start
func (s *AdminServer) LabelAdd(ctx context.Context, in *adminclient.LabelAddReq) (*adminclient.LabelAddResp, error) {
	l := logic.NewLabelAddLogic(ctx, s.svcCtx)
	return l.LabelAdd(in)
}

func (s *AdminServer) LabelList(ctx context.Context, in *adminclient.LabelListReq) (*adminclient.LabelListResp, error) {
	l := logic.NewLabelListLogic(ctx, s.svcCtx)
	return l.LabelList(in)
}

func (s *AdminServer) LabelUpdate(ctx context.Context, in *adminclient.LabelUpdateReq) (*adminclient.LabelUpdateResp, error) {
	l := logic.NewLabelUpdateLogic(ctx, s.svcCtx)
	return l.LabelUpdate(in)
}

func (s *AdminServer) LabelDelete(ctx context.Context, in *adminclient.LabelDeleteReq) (*adminclient.LabelDeleteResp, error) {
	l := logic.NewLabelDeleteLogic(ctx, s.svcCtx)
	return l.LabelDelete(in)
}

func (s *AdminServer) LabelListByPri(ctx context.Context, in *adminclient.LabelListByPriReq) (*adminclient.LabelListByPriResp, error) {
	l := logic.NewLabelListByPriLogic(ctx, s.svcCtx)
	return l.LabelListByPri(in)
}

// resource rpc start
func (s *AdminServer) ResourceAdd(ctx context.Context, in *adminclient.AddResourceReq) (*adminclient.AddResourceResp, error) {
	l := logic.NewResourceAddLogic(ctx, s.svcCtx)
	return l.ResourceAdd(in)
}

func (s *AdminServer) ResourceList(ctx context.Context, in *adminclient.ListResourceReq) (*adminclient.ListResourceResp, error) {
	l := logic.NewResourceListLogic(ctx, s.svcCtx)
	return l.ResourceList(in)
}

func (s *AdminServer) ResourceDelete(ctx context.Context, in *adminclient.DeleteResourceReq) (*adminclient.DeleteResourceResp, error) {
	l := logic.NewResourceDeleteLogic(ctx, s.svcCtx)
	return l.ResourceDelete(in)
}

func (s *AdminServer) ResourceObjectValueList(ctx context.Context, in *adminclient.ResourceObjectValueListReq) (*adminclient.ResourceObjectValueListResp, error) {
	l := logic.NewResourceObjectValueListLogic(ctx, s.svcCtx)
	return l.ResourceObjectValueList(in)
}

// search rpc start
func (s *AdminServer) SearchList(ctx context.Context, in *adminclient.SearchReq) (*adminclient.SearchResp, error) {
	l := logic.NewSearchListLogic(ctx, s.svcCtx)
	return l.SearchList(in)
}
