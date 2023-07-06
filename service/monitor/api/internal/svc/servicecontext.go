package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"ywadmin-v3/common/interceptor/metaserver"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/identity/rpc/identity"
	"ywadmin-v3/service/monitor/api/internal/config"
	"ywadmin-v3/service/monitor/rpc/monitor"
)

type ServiceContext struct {
	Config      config.Config
	Monitor     monitor.Monitor
	RedisClient *redis.Redis
	AdminRpc    admin.Admin
	IdentityRpc identity.Identity
}

func NewServiceContext(c config.Config) *ServiceContext {
	dialOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 1024))
	return &ServiceContext{
		Config:   c,
		Monitor:  monitor.NewMonitor(zrpc.MustNewClient(c.MonitorRpcConf, zrpc.WithUnaryClientInterceptor(metaserver.NameToInterceptor), zrpc.WithDialOption(dialOption))),
		AdminRpc: admin.NewAdmin(zrpc.MustNewClient(c.AdminRpcConf, zrpc.WithUnaryClientInterceptor(metaserver.NameToInterceptor), zrpc.WithDialOption(dialOption))),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		IdentityRpc: identity.NewIdentity(zrpc.MustNewClient(c.IdentityRpcConf)),
	}
}
