package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/techbloghub/server/config"
	"github.com/techbloghub/server/ent"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, cfgErr := config.NewConfig()
	if cfgErr != nil {
		log.Fatalf("failed to load config: %v", cfgErr)
	}

	pgCfg := cfg.PostgresConfig
	serverCfg := cfg.ServerConfig

	client, errPg := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		pgCfg.Host, pgCfg.Port, pgCfg.User, pgCfg.Db, pgCfg.Password))
	if errPg != nil {
		log.Fatalf("failed opening connection to postgres: %v", errPg)
	}
	defer client.Close()

	// Run the auto migration tool.
	if serverCfg.Env == "local" {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	// PORT 환경변수에서 가져오기
	// 후에 이런 config값 관리할것들 많아지면 후에 Config struct등으로 분리 고려
	r := setRouter()
	routerErr := r.Run(":" + serverCfg.Port)
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

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		fmt.Println("env get success", key, value)
		return value
	}
	fmt.Println("return default value for ", key)
	return defaultValue
}
