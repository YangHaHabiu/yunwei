package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//系统
	MonitorRpcConf  zrpc.RpcClientConf
	AdminRpcConf    zrpc.RpcClientConf
	IdentityRpcConf zrpc.RpcClientConf
	Redis           struct {
		Host string
		Type string
		Pass string
	}
	Auth struct {
		AccessSecret string
	}
}
