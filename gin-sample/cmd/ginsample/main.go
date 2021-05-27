package main

import (
	"gin-sample/cmd/ginsample/server"
	"gin-sample/conf"
)

func main() {
	cfg := conf.MustLoad()
	server.Run(cfg)
}
