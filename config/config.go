package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Static string
	Model  string
	DB     struct {
		Type   string
		Source string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
