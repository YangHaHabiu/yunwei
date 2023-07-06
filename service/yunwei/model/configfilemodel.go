package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ConfigFileModel = (*customConfigFileModel)(nil)

type (
	// ConfigFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConfigFileModel.
	ConfigFileModel interface {
		configFileModel
	}

	customConfigFileModel struct {
		*defaultConfigFileModel
	}
)

// NewConfigFileModel returns a model for the database table.
func NewConfigFileModel(conn sqlx.SqlConn) ConfigFileModel {
	return &customConfigFileModel{
		defaultConfigFileModel: newConfigFileModel(conn),
	}
}
