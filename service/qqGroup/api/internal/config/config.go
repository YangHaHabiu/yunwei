package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	Project struct {
		SupportGame      string
		YwTaskAddApiUrl  string
		HotUpdateApiUrl  string
		MaintainFilePath string
		InstallFilePath  string
		ManagerQQ        string
	}
}
