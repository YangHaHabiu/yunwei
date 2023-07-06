package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	IntranetRpcConf zrpc.RpcClientConf
	ScriptsFilePath string
	JwtAuth         struct {
		AccessSecret string
	}
}
