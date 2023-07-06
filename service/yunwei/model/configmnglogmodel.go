package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ConfigMngLogModel = (*customConfigMngLogModel)(nil)

type (
	// ConfigMngLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConfigMngLogModel.
	ConfigMngLogModel interface {
		configMngLogModel
	}

	customConfigMngLogModel struct {
		*defaultConfigMngLogModel
	}
)

// NewConfigMngLogModel returns a model for the database table.
func NewConfigMngLogModel(conn sqlx.SqlConn) ConfigMngLogModel {
	return &customConfigMngLogModel{
		defaultConfigMngLogModel: newConfigMngLogModel(conn),
	}
}
