package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProjectRelationshipModel = (*customProjectRelationshipModel)(nil)

type (
	// ProjectRelationshipModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectRelationshipModel.
	ProjectRelationshipModel interface {
		projectRelationshipModel
	}

	customProjectRelationshipModel struct {
		*defaultProjectRelationshipModel
	}
)

// NewProjectRelationshipModel returns a model for the database table.
func NewProjectRelationshipModel(conn sqlx.SqlConn) ProjectRelationshipModel {
	return &customProjectRelationshipModel{
		defaultProjectRelationshipModel: newProjectRelationshipModel(conn),
	}
}
