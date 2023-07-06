package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"ywadmin-v3/service/admin/model"
	"ywadmin-v3/service/admin/rpc/internal/config"
	"ywadmin-v3/service/identity/rpc/identity"
)

type ServiceContext struct {
	Config                   config.Config
	UserModel                model.SysUserModel
	MenuModel                model.SysMenuModel
	UserRoleModel            model.SysUserRoleModel
	UserProjectModel         model.SysUserProjectModel
	RoleModel                model.SysRoleModel
	RoleMenuModel            model.SysRoleMenuModel
	DictModel                model.SysDictModel
	DeptModel                model.SysDeptModel
	LoginLogModel            model.SysLoginLogModel
	SysLogModel              model.SysLogModel
	UserUgroupModel          model.SysUserUgroupModel
	UgroupModel              model.SysUgroupModel
	StrategyModel            model.SysStrategyModel
	StgroupUgroupModel       model.SysStgroupUgroupModel
	StgroupUserModel         model.SysStgroupUserModel
	StgroupModel             model.SysStgroupModel
	CompanyModel             model.CompanyModel
	ProjectModel             model.ProjectModel
	LabelModel               model.LabelModel
	LabelGlobalModel         model.LabelGlobalModel
	ProjectRelationshipModel model.ProjectRelationshipModel
	ServerAffiliationModel   model.ServerAffiliationModel
	RedisClient              *redis.Redis
	//rpc服务
	IdentityRpc identity.Identity
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		IdentityRpc:              identity.NewIdentity(zrpc.MustNewClient(c.IdentityRpcConf)),
		UserModel:                model.NewSysUserModel(sqlConn),
		MenuModel:                model.NewSysMenuModel(sqlConn),
		UserRoleModel:            model.NewSysUserRoleModel(sqlConn),
		UserProjectModel:         model.NewSysUserProjectModel(sqlConn),
		RoleModel:                model.NewSysRoleModel(sqlConn),
		RoleMenuModel:            model.NewSysRoleMenuModel(sqlConn),
		DictModel:                model.NewSysDictModel(sqlConn),
		DeptModel:                model.NewSysDeptModel(sqlConn),
		LoginLogModel:            model.NewSysLoginLogModel(sqlConn),
		SysLogModel:              model.NewSysLogModel(sqlConn),
		UserUgroupModel:          model.NewSysUserUgroupModel(sqlConn),
		UgroupModel:              model.NewSysUgroupModel(sqlConn),
		StrategyModel:            model.NewSysStrategyModel(sqlConn),
		StgroupUgroupModel:       model.NewSysStgroupUgroupModel(sqlConn),
		StgroupUserModel:         model.NewSysStgroupUserModel(sqlConn),
		StgroupModel:             model.NewSysStgroupModel(sqlConn),
		CompanyModel:             model.NewCompanyModel(sqlConn),
		ProjectModel:             model.NewProjectModel(sqlConn),
		LabelModel:               model.NewLabelModel(sqlConn),
		LabelGlobalModel:         model.NewLabelGlobalModel(sqlConn),
		ProjectRelationshipModel: model.NewProjectRelationshipModel(sqlConn),
		ServerAffiliationModel:   model.NewServerAffiliationModel(sqlConn),
	}
}
