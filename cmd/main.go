package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load(".env")
	if env != nil {
		return
	}
	r := SetRouter()
	err := r.Run(":" + os.Getenv("PORT"))
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
