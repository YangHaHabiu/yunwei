package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideVersionModel = (*customInsideVersionModel)(nil)

type (
	// InsideVersionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideVersionModel.
	InsideVersionModel interface {
		insideVersionModel
	}

	customInsideVersionModel struct {
		*defaultInsideVersionModel
	}
)

// NewInsideVersionModel returns a model for the database table.
func NewInsideVersionModel(conn sqlx.SqlConn) InsideVersionModel {
	return &customInsideVersionModel{
		defaultInsideVersionModel: newInsideVersionModel(conn),
	}
}
