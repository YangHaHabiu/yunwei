package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ LabelGlobalModel = (*customLabelGlobalModel)(nil)

type (
	// LabelGlobalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLabelGlobalModel.
	LabelGlobalModel interface {
		labelGlobalModel
	}

	customLabelGlobalModel struct {
		*defaultLabelGlobalModel
	}
)

// NewLabelGlobalModel returns a model for the database table.
func NewLabelGlobalModel(conn sqlx.SqlConn) LabelGlobalModel {
	return &customLabelGlobalModel{
		defaultLabelGlobalModel: newLabelGlobalModel(conn),
	}
}
