package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}

	AdminRpcConf   zrpc.RpcClientConf
	MonitorRpcConf zrpc.RpcClientConf

	TemplateFilePath string
	KeyFullPath      string
	ConfigCenterPath string
	LockFilePath     string
	ConfigMngThreads int
	IsOpenCall       bool
	Redis            struct {
		Host string
		Type string
		Pass string
	}
	Scripts struct {
		MaintainFilePath string
		InstallFilePath  string
		MigrateFilePath  string
		CombineFilePath  string
		FormatOutPutPath string
		InitScriptPath   string
	}
	YwQQGroup string
}
