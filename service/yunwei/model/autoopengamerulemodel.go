package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AutoOpengameRuleModel = (*customAutoOpengameRuleModel)(nil)

type (
	// AutoOpengameRuleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAutoOpengameRuleModel.
	AutoOpengameRuleModel interface {
		autoOpengameRuleModel
	}

	customAutoOpengameRuleModel struct {
		*defaultAutoOpengameRuleModel
	}
)

// NewAutoOpengameRuleModel returns a model for the database table.
func NewAutoOpengameRuleModel(conn sqlx.SqlConn) AutoOpengameRuleModel {
	return &customAutoOpengameRuleModel{
		defaultAutoOpengameRuleModel: newAutoOpengameRuleModel(conn),
	}
}
