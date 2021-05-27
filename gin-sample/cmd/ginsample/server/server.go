package server

import (
	"fmt"
	"os"
	"sync"

	"gin-sample/conf"
	"gin-sample/db"
	"gin-sample/global"
	"gin-sample/handler"
	"gin-sample/router"
)

var once sync.Once

func Run(cfg *conf.Config) {
	onceInit(cfg)
	engine := router.Init()

	if err := engine.Run(cfg.ApiCfg.APIAddr); err != nil {
		fmt.Printf("start server error: %v", err.Error())
		os.Exit(1)
	}
}

func onceInit(cfg *conf.Config) {
	once.Do(func() {
		var err error
		if err = global.Init(cfg.GlobalCfg); err != nil {
			fmt.Printf("Init global error: %v", err.Error())
			os.Exit(1)
		}

		if err = handler.Init(); err != nil {
			fmt.Printf("Init handler error: %v", err.Error())
			os.Exit(1)
		}

		if err = db.Init(cfg.DBCfg); err != nil {
			fmt.Printf("Init Database error: %v", err.Error())
			os.Exit(1)
		}
	})
}
