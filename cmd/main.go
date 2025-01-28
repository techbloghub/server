package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/techbloghub/server/config"
	"github.com/techbloghub/server/ent"
	_ "github.com/techbloghub/server/ent/runtime"
	"github.com/techbloghub/server/internal/database"
	"github.com/techbloghub/server/internal/http/router"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Print("failed to reading .env", errEnv)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r, dbClient, err := createServer(cfg)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	defer dbClient.Close()

	if err := r.Run(":" + cfg.ServerConfig.Port); err != nil {
		log.Fatalf("Error while running server: %v", err)
	}
}

func createServer(cfg *config.Config) (*gin.Engine, *ent.Client, error) {
	client, err := database.ConnectDatabase(cfg)
	if err != nil {
		return nil, nil, err
	}

	r := gin.Default()
	router.InitRouter(r, client)

	return r, client, nil
}
