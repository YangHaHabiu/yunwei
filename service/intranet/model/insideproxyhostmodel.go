package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideProxyHostModel = (*customInsideProxyHostModel)(nil)

type (
	// InsideProxyHostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideProxyHostModel.
	InsideProxyHostModel interface {
		insideProxyHostModel
	}

	customInsideProxyHostModel struct {
		*defaultInsideProxyHostModel
	}
)

// NewInsideProxyHostModel returns a model for the database table.
func NewInsideProxyHostModel(conn sqlx.SqlConn) InsideProxyHostModel {
	return &customInsideProxyHostModel{
		defaultInsideProxyHostModel: newInsideProxyHostModel(conn),
	}
}
