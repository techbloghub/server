package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/techbloghub/server/ent"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Print("failed to reading .env", errEnv)
	}

	postgresHost := getEnvWithDefault("POSTGRES_HOST", "localhost")
	postgresPort := getEnvWithDefault("POSTGRES_PORT", "5432")
	postgresUser := getEnvWithDefault("POSTGRES_USER", "tbh-user")
	postgresPassword := getEnvWithDefault("POSTGRES_PASSWORD", "password")
	postgresDbname := getEnvWithDefault("POSTGRES_DB", "tbh-db")

	client, errPg := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresDbname, postgresPassword))
	if errPg != nil {
		log.Fatalf("failed opening connection to postgres: %v", errPg)
	}
	defer client.Close()
	// Run the auto migration tool.
	if env := getEnvWithDefault("ENV", "local"); env == "local" {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

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
		fmt.Println("env get success", key, value)
		return value
	}
	fmt.Println("return default value for ", key)
	return defaultValue
}
