package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TasksTidPidModel = (*customTasksTidPidModel)(nil)

type (
	// TasksTidPidModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTasksTidPidModel.
	TasksTidPidModel interface {
		tasksTidPidModel
	}

	customTasksTidPidModel struct {
		*defaultTasksTidPidModel
	}
)

// NewTasksTidPidModel returns a model for the database table.
func NewTasksTidPidModel(conn sqlx.SqlConn) TasksTidPidModel {
	return &customTasksTidPidModel{
		defaultTasksTidPidModel: newTasksTidPidModel(conn),
	}
}
