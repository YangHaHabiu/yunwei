package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	YunWeiRpcConf zrpc.RpcClientConf
	AdminRpcConf  zrpc.RpcClientConf
	JwtAuth       struct {
		AccessSecret string
	}
	Scripts struct {
		MaintainFilePath string
		InstallFilePath  string
		MigrateFilePath  string
		CombineFilePath  string
		FormatOutPutPath string
	}

	JumpServer []struct {
		Ipaddr         string
		PrivateKeyPath string
		SocksName      string
		SocksPwd       string
	}
}
