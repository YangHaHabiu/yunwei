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
	YunWeiRpcConf zrpc.RpcClientConf
	AdminRpcConf  zrpc.RpcClientConf
	Project       struct {
		SupportGame      string
		MaintainFilePath string
		InstallFilePath  string
		ManagerQQ        string
		LockFilePath     string
	}
	YwQQGroup      string
	IsOpenCheckVpn bool
}
