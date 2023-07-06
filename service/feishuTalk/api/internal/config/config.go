package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	AppId             string
	AppSecret         string
	EncryptKey        string
	VerificationToken string
}
