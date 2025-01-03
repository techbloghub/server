package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// PORT 환경변수에서 가져오기
	// 후에 이런 config값 관리할것들 많아지면 후에 Config struct등으로 분리 고려
	port := getEnvWithDefault("PORT", "8080")
	r := setRouter()
	err := r.Run(":" + port)
	if err != nil {
		fmt.Println("Error while running server: ", err)
		return
	}
}

func setRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
	return r
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
