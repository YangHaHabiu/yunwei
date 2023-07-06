package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideServerModel = (*customInsideServerModel)(nil)

type (
	// InsideServerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideServerModel.
	InsideServerModel interface {
		insideServerModel
	}

	customInsideServerModel struct {
		*defaultInsideServerModel
	}
)

// NewInsideServerModel returns a model for the database table.
func NewInsideServerModel(conn sqlx.SqlConn) InsideServerModel {
	return &customInsideServerModel{
		defaultInsideServerModel: newInsideServerModel(conn),
	}
}
