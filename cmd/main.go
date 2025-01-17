package main

import (
	"fmt"
	"log"

	"github.com/techbloghub/server/config"
	"github.com/techbloghub/server/internal/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, cfgErr := config.NewConfig()
	if cfgErr != nil {
		log.Fatalf("failed to load config: %v", cfgErr)
	}

	// DB 연결
	client, errPg := database.ConnectDatabase(cfg)
	if errPg != nil {
		log.Fatalf("failed to connect database: %v", errPg)
	}
	defer client.Close()

	// 서버 실행
	r := setRouter()
	routerErr := r.Run(":" + cfg.ServerConfig.Port)
	if routerErr != nil {
		fmt.Println("Error while running server: ", cfgErr)
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
