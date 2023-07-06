package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideTasksModel = (*customInsideTasksModel)(nil)

type (
	// InsideTasksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideTasksModel.
	InsideTasksModel interface {
		insideTasksModel
	}

	customInsideTasksModel struct {
		*defaultInsideTasksModel
	}
)

// NewInsideTasksModel returns a model for the database table.
func NewInsideTasksModel(conn sqlx.SqlConn) InsideTasksModel {
	return &customInsideTasksModel{
		defaultInsideTasksModel: newInsideTasksModel(conn),
	}
}
