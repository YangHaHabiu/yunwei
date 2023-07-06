package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideInstallPlanModel = (*customInsideInstallPlanModel)(nil)

type (
	// InsideInstallPlanModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideInstallPlanModel.
	InsideInstallPlanModel interface {
		insideInstallPlanModel
	}

	customInsideInstallPlanModel struct {
		*defaultInsideInstallPlanModel
	}
)

// NewInsideInstallPlanModel returns a model for the database table.
func NewInsideInstallPlanModel(conn sqlx.SqlConn) InsideInstallPlanModel {
	return &customInsideInstallPlanModel{
		defaultInsideInstallPlanModel: newInsideInstallPlanModel(conn),
	}
}
