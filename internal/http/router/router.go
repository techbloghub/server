package router

import (
	"github.com/gin-gonic/gin"
	"github.com/techbloghub/server/internal/http/handler"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", handler.PingPong)
}
