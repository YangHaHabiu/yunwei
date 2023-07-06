package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideOperationModel = (*customInsideOperationModel)(nil)

type (
	// InsideOperationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideOperationModel.
	InsideOperationModel interface {
		insideOperationModel
	}

	customInsideOperationModel struct {
		*defaultInsideOperationModel
	}
)

// NewInsideOperationModel returns a model for the database table.
func NewInsideOperationModel(conn sqlx.SqlConn) InsideOperationModel {
	return &customInsideOperationModel{
		defaultInsideOperationModel: newInsideOperationModel(conn),
	}
}
