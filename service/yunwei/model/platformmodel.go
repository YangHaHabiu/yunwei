package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PlatformModel = (*customPlatformModel)(nil)

type (
	// PlatformModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlatformModel.
	PlatformModel interface {
		platformModel
	}

	customPlatformModel struct {
		*defaultPlatformModel
	}
)

// NewPlatformModel returns a model for the database table.
func NewPlatformModel(conn sqlx.SqlConn) PlatformModel {
	return &customPlatformModel{
		defaultPlatformModel: newPlatformModel(conn),
	}
}
