package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TaskLogHistroyModel = (*customTaskLogHistroyModel)(nil)

type (
	// TaskLogHistroyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskLogHistroyModel.
	TaskLogHistroyModel interface {
		taskLogHistroyModel
	}

	customTaskLogHistroyModel struct {
		*defaultTaskLogHistroyModel
	}
)

// NewTaskLogHistroyModel returns a model for the database table.
func NewTaskLogHistroyModel(conn sqlx.SqlConn) TaskLogHistroyModel {
	return &customTaskLogHistroyModel{
		defaultTaskLogHistroyModel: newTaskLogHistroyModel(conn),
	}
}
