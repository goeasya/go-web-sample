package router

import (
	"net/http"

	"gin-sample/api"
	"gin-sample/middleware"
	"gin-sample/util"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() *gin.Engine {
	// 设置日志文件切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "/var/log/goweb/gin.log",
		MaxSize:    100,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
	}
	gin.DefaultWriter = lumberJackLogger
	return SetupRouter()
}

func SetupRouter() *gin.Engine {

	r := gin.Default()
	pprof.Register(r)
	r.Use(middleware.CorsHandler())
	r.GET("/health", Health)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/user/register", api.Register)
	r.POST("/user/login", api.Login)

	// api
	apiR := r.Group("api")
	apiR.Use(middleware.JWT())
	InitV1Routes(apiR)

	return r
}

func Health(c *gin.Context) {
	util.ResponseSuccess(c, http.StatusOK, "gin sample forever!")
}
