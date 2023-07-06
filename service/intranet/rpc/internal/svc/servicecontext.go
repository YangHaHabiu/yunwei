package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"ywadmin-v3/common/interceptor/metaserver"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/intranet/model"
	"ywadmin-v3/service/intranet/rpc/internal/config"
	"ywadmin-v3/service/yunwei/rpc/yunwei"
)

type ServiceContext struct {
	Config                    config.Config
	RedisClient               *redis.Redis
	AdminRpc                  admin.Admin
	YunWeiRpc                 yunwei.YunWei
	InsideInstallPlanModel    model.InsideInstallPlanModel
	InsideOperationModel      model.InsideOperationModel
	InsideProjectClusterModel model.InsideProjectClusterModel
	InsideServerModel         model.InsideServerModel
	InsideTasksLogsModel      model.InsideTasksLogsModel
	InsideTasksModel          model.InsideTasksModel
	InsideTasksPidModel       model.InsideTasksPidModel
	InsideProxyHostModel      model.InsideProxyHostModel
	InsideVersionModel        model.InsideVersionModel
	InsideHostInfoModel       model.InsideHostInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dialOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 1024))
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		YunWeiRpc: yunwei.NewYunWei(zrpc.MustNewClient(c.YunWeiRpcConf, zrpc.WithUnaryClientInterceptor(metaserver.NameToInterceptor), zrpc.WithDialOption(dialOption))),
		AdminRpc:  admin.NewAdmin(zrpc.MustNewClient(c.AdminRpcConf, zrpc.WithUnaryClientInterceptor(metaserver.NameToInterceptor), zrpc.WithDialOption(dialOption))),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		InsideInstallPlanModel:    model.NewInsideInstallPlanModel(sqlConn),
		InsideOperationModel:      model.NewInsideOperationModel(sqlConn),
		InsideProjectClusterModel: model.NewInsideProjectClusterModel(sqlConn),
		InsideServerModel:         model.NewInsideServerModel(sqlConn),
		InsideTasksLogsModel:      model.NewInsideTasksLogsModel(sqlConn),
		InsideTasksModel:          model.NewInsideTasksModel(sqlConn),
		InsideTasksPidModel:       model.NewInsideTasksPidModel(sqlConn),
		InsideProxyHostModel:      model.NewInsideProxyHostModel(sqlConn),
		InsideVersionModel:        model.NewInsideVersionModel(sqlConn),
		InsideHostInfoModel:       model.NewInsideHostInfoModel(sqlConn),
	}
}
