package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"ywadmin-v3/service/monitor/model"
	"ywadmin-v3/service/monitor/rpc/internal/config"
)

type ServiceContext struct {
	Config                  config.Config
	ReportStreamMinuteModel model.ReportStreamMinuteModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                  c,
		ReportStreamMinuteModel: model.NewReportStreamMinuteModel(sqlConn),
	}
}
