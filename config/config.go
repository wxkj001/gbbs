package config

import "github.com/zeromicro/go-zero/rest"

type User struct {
	VisitorUUID  string
	OrdinaryUUID string
	AdminUUID    string
}
type Config struct {
	rest.RestConf
	Debug  bool
	Static string
	Rbac   string
	Pass   string
	User   User
	DB     struct {
		Type   string
		Source string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
