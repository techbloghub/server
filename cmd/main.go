package main

import (
	"fmt"
	"log"

	"github.com/techbloghub/server/config"
	_ "github.com/techbloghub/server/ent/runtime"
	"github.com/techbloghub/server/internal/database"
	"github.com/techbloghub/server/internal/http/router"

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
	r := gin.Default()
	router.InitRouter(r)
	routerErr := r.Run(":" + cfg.ServerConfig.Port)
	if routerErr != nil {
		fmt.Println("Error while running server: ", routerErr)
		return
	}
}
