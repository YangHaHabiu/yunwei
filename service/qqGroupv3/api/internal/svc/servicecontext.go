package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/qqGroupv3/api/internal/config"
	"ywadmin-v3/service/qqGroupv3/model"
	"ywadmin-v3/service/yunwei/rpc/yunwei"
)

type ServiceContext struct {
	Config                config.Config
	QqMessageHistoryModel model.QqMessageHistoryModel
	QqLoadBalanceModel    model.QqLoadBalanceModel
	YunWeiRpc             yunwei.YunWei
	AdminRpc              admin.Admin
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                c,
		QqLoadBalanceModel:    model.NewQqLoadBalanceModel(sqlConn),
		QqMessageHistoryModel: model.NewQqMessageHistoryModel(sqlConn),
		AdminRpc:              admin.NewAdmin(zrpc.MustNewClient(c.AdminRpcConf)),
		YunWeiRpc:             yunwei.NewYunWei(zrpc.MustNewClient(c.YunWeiRpcConf)),
	}
}
