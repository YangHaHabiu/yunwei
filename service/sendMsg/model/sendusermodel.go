package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SendUserModel = (*customSendUserModel)(nil)

type (
	// SendUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSendUserModel.
	SendUserModel interface {
		sendUserModel
	}

	customSendUserModel struct {
		*defaultSendUserModel
	}
)

// NewSendUserModel returns a model for the database table.
func NewSendUserModel(conn sqlx.SqlConn) SendUserModel {
	return &customSendUserModel{
		defaultSendUserModel: newSendUserModel(conn),
	}
}
