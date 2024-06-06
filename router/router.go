package router

import (
	"20240606/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.Roothandler)

	r.POST("/POST", handler.Posthandler)
}
