package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"ywadmin-v3/service/admin/rpc/admin"
	"ywadmin-v3/service/ws/api/internal/config"
	"ywadmin-v3/service/yunwei/rpc/yunwei"
)

type ServiceContext struct {
	Config    config.Config
	YunWeiRpc yunwei.YunWei
	AdminRpc  admin.Admin
}

func NewServiceContext(c config.Config) *ServiceContext {
	dialOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 1024))
	return &ServiceContext{
		Config:    c,
		YunWeiRpc: yunwei.NewYunWei(zrpc.MustNewClient(c.YunWeiRpcConf, zrpc.WithDialOption(dialOption))),
		AdminRpc:  admin.NewAdmin(zrpc.MustNewClient(c.AdminRpcConf, zrpc.WithDialOption(dialOption))),
	}
}
