package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideTasksPidModel = (*customInsideTasksPidModel)(nil)

type (
	// InsideTasksPidModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideTasksPidModel.
	InsideTasksPidModel interface {
		insideTasksPidModel
	}

	customInsideTasksPidModel struct {
		*defaultInsideTasksPidModel
	}
)

// NewInsideTasksPidModel returns a model for the database table.
func NewInsideTasksPidModel(conn sqlx.SqlConn) InsideTasksPidModel {
	return &customInsideTasksPidModel{
		defaultInsideTasksPidModel: newInsideTasksPidModel(conn),
	}
}
