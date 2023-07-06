package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ReportStreamMinuteModel = (*customReportStreamMinuteModel)(nil)

type (
	// ReportStreamMinuteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReportStreamMinuteModel.
	ReportStreamMinuteModel interface {
		reportStreamMinuteModel
	}

	customReportStreamMinuteModel struct {
		*defaultReportStreamMinuteModel
	}
)

// NewReportStreamMinuteModel returns a model for the database table.
func NewReportStreamMinuteModel(conn sqlx.SqlConn) ReportStreamMinuteModel {
	return &customReportStreamMinuteModel{
		defaultReportStreamMinuteModel: newReportStreamMinuteModel(conn),
	}
}
