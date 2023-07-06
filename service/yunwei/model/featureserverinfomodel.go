package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FeatureServerInfoModel = (*customFeatureServerInfoModel)(nil)

type (
	// FeatureServerInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFeatureServerInfoModel.
	FeatureServerInfoModel interface {
		featureServerInfoModel
	}

	customFeatureServerInfoModel struct {
		*defaultFeatureServerInfoModel
	}
)

// NewFeatureServerInfoModel returns a model for the database table.
func NewFeatureServerInfoModel(conn sqlx.SqlConn) FeatureServerInfoModel {
	return &customFeatureServerInfoModel{
		defaultFeatureServerInfoModel: newFeatureServerInfoModel(conn),
	}
}
