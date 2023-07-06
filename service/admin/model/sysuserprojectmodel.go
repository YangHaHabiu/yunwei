package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysUserProjectModel = (*customSysUserProjectModel)(nil)

type (
	// SysUserProjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserProjectModel.
	SysUserProjectModel interface {
		sysUserProjectModel
	}

	customSysUserProjectModel struct {
		*defaultSysUserProjectModel
	}
)

// NewSysUserProjectModel returns a model for the database table.
func NewSysUserProjectModel(conn sqlx.SqlConn) SysUserProjectModel {
	return &customSysUserProjectModel{
		defaultSysUserProjectModel: newSysUserProjectModel(conn),
	}
}
