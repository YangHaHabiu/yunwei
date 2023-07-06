package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"ywadmin-v3/common/interceptor/metaserver"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/identity/rpc/identity"
	"ywadmin-v3/service/sendMsg/api/internal/config"
	"ywadmin-v3/service/sendMsg/model"
)

type ServiceContext struct {
	Config             config.Config
	RedisClient        *redis.Redis
	AdminRpc           admin.Admin
	IdentityRpc        identity.Identity
	SendMsgRecordModel model.SendMsgRecordModel
	SendAccountModel   model.SendAccountModel
	SendUserModel      model.SendUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dialOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 1024))
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		AdminRpc:    admin.NewAdmin(zrpc.MustNewClient(c.AdminRpcConf, zrpc.WithUnaryClientInterceptor(metaserver.NameToInterceptor), zrpc.WithDialOption(dialOption))),
		IdentityRpc: identity.NewIdentity(zrpc.MustNewClient(c.IdentityRpcConf, zrpc.WithUnaryClientInterceptor(metaserver.NameToInterceptor), zrpc.WithDialOption(dialOption))),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		SendMsgRecordModel: model.NewSendMsgRecordModel(sqlConn),
		SendAccountModel:   model.NewSendAccountModel(sqlConn),
		SendUserModel:      model.NewSendUserModel(sqlConn),
	}
}
