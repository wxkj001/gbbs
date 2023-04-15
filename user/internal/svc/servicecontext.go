package svc

import (
	"bbs/config"
	"bbs/ent"
)

type ServiceContext struct {
	Config config.Config
	Ent    *ent.Client
}

func NewServiceContext(c config.Config, db *ent.Client) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Ent:    db,
	}
}
