package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	ApiAuthKey struct {
		List []struct {
			Name      string
			AppKey    string
			AppSecret string
		}
		Limit          int64
		DefaultChannel string
		CardSending    bool
	}
	//系统
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
