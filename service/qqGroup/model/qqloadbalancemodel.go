package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ QqLoadBalanceModel = (*customQqLoadBalanceModel)(nil)

type (
	// QqLoadBalanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQqLoadBalanceModel.
	QqLoadBalanceModel interface {
		qqLoadBalanceModel
	}

	customQqLoadBalanceModel struct {
		*defaultQqLoadBalanceModel
	}
)

// NewQqLoadBalanceModel returns a model for the database table.
func NewQqLoadBalanceModel(conn sqlx.SqlConn) QqLoadBalanceModel {
	return &customQqLoadBalanceModel{
		defaultQqLoadBalanceModel: newQqLoadBalanceModel(conn),
	}
}
