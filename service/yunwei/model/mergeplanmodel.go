package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MergePlanModel = (*customMergePlanModel)(nil)

type (
	// MergePlanModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMergePlanModel.
	MergePlanModel interface {
		mergePlanModel
	}

	customMergePlanModel struct {
		*defaultMergePlanModel
	}
)

// NewMergePlanModel returns a model for the database table.
func NewMergePlanModel(conn sqlx.SqlConn) MergePlanModel {
	return &customMergePlanModel{
		defaultMergePlanModel: newMergePlanModel(conn),
	}
}
