package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AlarmThresholdManageModel = (*customAlarmThresholdManageModel)(nil)

type (
	// AlarmThresholdManageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAlarmThresholdManageModel.
	AlarmThresholdManageModel interface {
		alarmThresholdManageModel
	}

	customAlarmThresholdManageModel struct {
		*defaultAlarmThresholdManageModel
	}
)

// NewAlarmThresholdManageModel returns a model for the database table.
func NewAlarmThresholdManageModel(conn sqlx.SqlConn) AlarmThresholdManageModel {
	return &customAlarmThresholdManageModel{
		defaultAlarmThresholdManageModel: newAlarmThresholdManageModel(conn),
	}
}
