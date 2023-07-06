package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	//系统
	AdminRpcConf  zrpc.RpcClientConf
	YunWeiRpcConf zrpc.RpcClientConf
	DB            struct {
		DataSource string
	}
	Redis struct {
		Host string
		Type string
		Pass string
	}
}
