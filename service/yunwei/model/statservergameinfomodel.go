package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StatServerGameInfoModel = (*customStatServerGameInfoModel)(nil)

type (
	// StatServerGameInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStatServerGameInfoModel.
	StatServerGameInfoModel interface {
		statServerGameInfoModel
	}

	customStatServerGameInfoModel struct {
		*defaultStatServerGameInfoModel
	}
)

// NewStatServerGameInfoModel returns a model for the database table.
func NewStatServerGameInfoModel(conn sqlx.SqlConn) StatServerGameInfoModel {
	return &customStatServerGameInfoModel{
		defaultStatServerGameInfoModel: newStatServerGameInfoModel(conn),
	}
}
