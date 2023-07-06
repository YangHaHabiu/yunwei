package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideHostInfoModel = (*customInsideHostInfoModel)(nil)

type (
	// InsideHostInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideHostInfoModel.
	InsideHostInfoModel interface {
		insideHostInfoModel
	}

	customInsideHostInfoModel struct {
		*defaultInsideHostInfoModel
	}
)

// NewInsideHostInfoModel returns a model for the database table.
func NewInsideHostInfoModel(conn sqlx.SqlConn) InsideHostInfoModel {
	return &customInsideHostInfoModel{
		defaultInsideHostInfoModel: newInsideHostInfoModel(conn),
	}
}
