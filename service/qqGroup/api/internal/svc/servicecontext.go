package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"ywadmin-v3/service/qqGroup/api/internal/config"
	"ywadmin-v3/service/qqGroup/model"
)

type ServiceContext struct {
	Config                config.Config
	QqMessageHistoryModel model.QqMessageHistoryModel
	QqLoadBalanceModel    model.QqLoadBalanceModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                c,
		QqLoadBalanceModel:    model.NewQqLoadBalanceModel(sqlConn),
		QqMessageHistoryModel: model.NewQqMessageHistoryModel(sqlConn),
	}
}
