package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MaintainPlanModel = (*customMaintainPlanModel)(nil)

type (
	// MaintainPlanModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMaintainPlanModel.
	MaintainPlanModel interface {
		maintainPlanModel
	}

	customMaintainPlanModel struct {
		*defaultMaintainPlanModel
	}
)

// NewMaintainPlanModel returns a model for the database table.
func NewMaintainPlanModel(conn sqlx.SqlConn) MaintainPlanModel {
	return &customMaintainPlanModel{
		defaultMaintainPlanModel: newMaintainPlanModel(conn),
	}
}
