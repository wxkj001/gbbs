package main

import (
	"bbs/config"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	app, cleanup, err := InitializeEvent(c)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	app.Run()
}
