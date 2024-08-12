package main

import (
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	r := SetRouter()
	err := r.Run(":" + PORT)
	if err != nil {
		return
	}
}

func SetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
	return r
}
