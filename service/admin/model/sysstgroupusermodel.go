package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysStgroupUserModel = (*customSysStgroupUserModel)(nil)

type (
	// SysStgroupUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysStgroupUserModel.
	SysStgroupUserModel interface {
		sysStgroupUserModel
	}

	customSysStgroupUserModel struct {
		*defaultSysStgroupUserModel
	}
)

// NewSysStgroupUserModel returns a model for the database table.
func NewSysStgroupUserModel(conn sqlx.SqlConn) SysStgroupUserModel {
	return &customSysStgroupUserModel{
		defaultSysStgroupUserModel: newSysStgroupUserModel(conn),
	}
}
