package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"ywadmin-v3/service/intranet/rpc/intranet"
	"ywadmin-v3/service/inws/api/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	IntranetRpc intranet.Intranet
}

func NewServiceContext(c config.Config) *ServiceContext {
	dialOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 1024))
	return &ServiceContext{
		Config:      c,
		IntranetRpc: intranet.NewIntranet(zrpc.MustNewClient(c.IntranetRpcConf, zrpc.WithDialOption(dialOption))),
	}
}
