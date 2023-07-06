package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AssetModel = (*customAssetModel)(nil)

type (
	// AssetModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAssetModel.
	AssetModel interface {
		assetModel
	}

	customAssetModel struct {
		*defaultAssetModel
	}
)

// NewAssetModel returns a model for the database table.
func NewAssetModel(conn sqlx.SqlConn) AssetModel {
	return &customAssetModel{
		defaultAssetModel: newAssetModel(conn),
	}
}
