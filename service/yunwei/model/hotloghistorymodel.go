package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HotLogHistoryModel = (*customHotLogHistoryModel)(nil)

type (
	// HotLogHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHotLogHistoryModel.
	HotLogHistoryModel interface {
		hotLogHistoryModel
	}

	customHotLogHistoryModel struct {
		*defaultHotLogHistoryModel
	}
)

// NewHotLogHistoryModel returns a model for the database table.
func NewHotLogHistoryModel(conn sqlx.SqlConn) HotLogHistoryModel {
	return &customHotLogHistoryModel{
		defaultHotLogHistoryModel: newHotLogHistoryModel(conn),
	}
}
