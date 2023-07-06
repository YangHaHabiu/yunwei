package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProjectModel = (*customProjectModel)(nil)

type (
	// ProjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectModel.
	ProjectModel interface {
		projectModel
	}

	customProjectModel struct {
		*defaultProjectModel
	}
)

// NewProjectModel returns a model for the database table.
func NewProjectModel(conn sqlx.SqlConn) ProjectModel {
	return &customProjectModel{
		defaultProjectModel: newProjectModel(conn),
	}
}
