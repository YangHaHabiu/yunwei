package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OpenPlanModel = (*customOpenPlanModel)(nil)

type (
	// OpenPlanModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOpenPlanModel.
	OpenPlanModel interface {
		openPlanModel
	}

	customOpenPlanModel struct {
		*defaultOpenPlanModel
	}
)

// NewOpenPlanModel returns a model for the database table.
func NewOpenPlanModel(conn sqlx.SqlConn) OpenPlanModel {
	return &customOpenPlanModel{
		defaultOpenPlanModel: newOpenPlanModel(conn),
	}
}
