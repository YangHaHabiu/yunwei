package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	IdentityRpcConf zrpc.RpcClientConf
	Redis           struct {
		Host string
		Type string
		Pass string
	}
}
