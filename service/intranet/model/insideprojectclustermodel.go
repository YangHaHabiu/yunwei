package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InsideProjectClusterModel = (*customInsideProjectClusterModel)(nil)

type (
	// InsideProjectClusterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInsideProjectClusterModel.
	InsideProjectClusterModel interface {
		insideProjectClusterModel
	}

	customInsideProjectClusterModel struct {
		*defaultInsideProjectClusterModel
	}
)

// NewInsideProjectClusterModel returns a model for the database table.
func NewInsideProjectClusterModel(conn sqlx.SqlConn) InsideProjectClusterModel {
	return &customInsideProjectClusterModel{
		defaultInsideProjectClusterModel: newInsideProjectClusterModel(conn),
	}
}
