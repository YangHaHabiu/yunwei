package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ QqMessageHistoryModel = (*customQqMessageHistoryModel)(nil)

type (
	// QqMessageHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQqMessageHistoryModel.
	QqMessageHistoryModel interface {
		qqMessageHistoryModel
	}

	customQqMessageHistoryModel struct {
		*defaultQqMessageHistoryModel
	}
)

// NewQqMessageHistoryModel returns a model for the database table.
func NewQqMessageHistoryModel(conn sqlx.SqlConn) QqMessageHistoryModel {
	return &customQqMessageHistoryModel{
		defaultQqMessageHistoryModel: newQqMessageHistoryModel(conn),
	}
}
