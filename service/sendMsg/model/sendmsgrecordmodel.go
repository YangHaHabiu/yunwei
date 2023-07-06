package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SendMsgRecordModel = (*customSendMsgRecordModel)(nil)

type (
	// SendMsgRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSendMsgRecordModel.
	SendMsgRecordModel interface {
		sendMsgRecordModel
	}

	customSendMsgRecordModel struct {
		*defaultSendMsgRecordModel
	}
)

// NewSendMsgRecordModel returns a model for the database table.
func NewSendMsgRecordModel(conn sqlx.SqlConn) SendMsgRecordModel {
	return &customSendMsgRecordModel{
		defaultSendMsgRecordModel: newSendMsgRecordModel(conn),
	}
}
