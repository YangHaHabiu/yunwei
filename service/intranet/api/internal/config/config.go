package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//系统
	AdminRpcConf    zrpc.RpcClientConf
	YunWeiRpcConf   zrpc.RpcClientConf
	IdentityRpcConf zrpc.RpcClientConf
	IntranetRpcConf zrpc.RpcClientConf
	Redis           struct {
		Host string
		Type string
		Pass string
	}
	Auth struct {
		AccessSecret string
	}
}
