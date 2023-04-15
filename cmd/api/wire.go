//go:build wireinject
// +build wireinject

// wire.go
package main

import (
	"bbs/app"
	"bbs/config"
	"github.com/google/wire"
)

func InitializeEvent(c config.Config) (*app.App, func(), error) {
	panic(wire.Build(app.NewServer, app.NewApp, app.NewDB, app.NewCasbin))
}
