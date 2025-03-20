package handler

import "github.com/gin-gonic/gin"

func PingPong(context *gin.Context) {
	context.String(200, "pong\n")
}

func DoNothing(context *gin.Context) {
	context.Status(200)
}
