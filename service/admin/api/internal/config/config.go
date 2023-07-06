package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//系统
	AdminRpcConf              zrpc.RpcClientConf
	IdentityRpcConf           zrpc.RpcClientConf
	VerificationCode          bool
	VerificationCodeWatermark string
	Path                      string `json:",default=./files"`
	Redis                     struct {
		Host string
		Type string
		Pass string
	}
	Scripts struct {
		MaintainFilePath string
	}
	RecPath struct {
		FullPath string
	}
	Auth struct {
		AccessSecret string
	}
}
