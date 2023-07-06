package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ KeyManageModel = (*customKeyManageModel)(nil)

type (
	// KeyManageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customKeyManageModel.
	KeyManageModel interface {
		keyManageModel
	}

	customKeyManageModel struct {
		*defaultKeyManageModel
	}
)

// NewKeyManageModel returns a model for the database table.
func NewKeyManageModel(conn sqlx.SqlConn) KeyManageModel {
	return &customKeyManageModel{
		defaultKeyManageModel: newKeyManageModel(conn),
	}
}
