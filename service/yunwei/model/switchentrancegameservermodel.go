package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SwitchEntranceGameserverModel = (*customSwitchEntranceGameserverModel)(nil)

type (
	// SwitchEntranceGameserverModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSwitchEntranceGameserverModel.
	SwitchEntranceGameserverModel interface {
		switchEntranceGameserverModel
	}

	customSwitchEntranceGameserverModel struct {
		*defaultSwitchEntranceGameserverModel
	}
)

// NewSwitchEntranceGameserverModel returns a model for the database table.
func NewSwitchEntranceGameserverModel(conn sqlx.SqlConn) SwitchEntranceGameserverModel {
	return &customSwitchEntranceGameserverModel{
		defaultSwitchEntranceGameserverModel: newSwitchEntranceGameserverModel(conn),
	}
}
