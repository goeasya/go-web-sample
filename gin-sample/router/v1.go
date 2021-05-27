package router

import (
	v1 "gin-sample/api/v1"

	"github.com/gin-gonic/gin"
)

func InitV1Routes(rg *gin.RouterGroup) {
	r := rg.Group("/v1")
	initUserRoutes(r)
}

func initUserRoutes(rg *gin.RouterGroup) {
	userR := rg.Group("/user")
	userR.DELETE("", v1.DeleteUser)
	userR.PUT("", v1.UpdateUser)
	userR.GET("/detail", v1.GetUserDetail)
}
