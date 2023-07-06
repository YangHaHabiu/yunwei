package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ServerAffiliationModel = (*customServerAffiliationModel)(nil)

type (
	// ServerAffiliationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customServerAffiliationModel.
	ServerAffiliationModel interface {
		serverAffiliationModel
	}

	customServerAffiliationModel struct {
		*defaultServerAffiliationModel
	}
)

// NewServerAffiliationModel returns a model for the database table.
func NewServerAffiliationModel(conn sqlx.SqlConn) ServerAffiliationModel {
	return &customServerAffiliationModel{
		defaultServerAffiliationModel: newServerAffiliationModel(conn),
	}
}
