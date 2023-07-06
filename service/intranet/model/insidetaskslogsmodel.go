package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideTasksLogsModel = (*customInsideTasksLogsModel)(nil)

type (
	// InsideTasksLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideTasksLogsModel.
	InsideTasksLogsModel interface {
		insideTasksLogsModel
	}

	customInsideTasksLogsModel struct {
		*defaultInsideTasksLogsModel
	}
)

// NewInsideTasksLogsModel returns a model for the database table.
func NewInsideTasksLogsModel(conn sqlx.SqlConn) InsideTasksLogsModel {
	return &customInsideTasksLogsModel{
		defaultInsideTasksLogsModel: newInsideTasksLogsModel(conn),
	}
}
