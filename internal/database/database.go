package database

import (
	"context"
	"fmt"

	"github.com/techbloghub/server/config"
	"github.com/techbloghub/server/ent"
)

func ConnectDatabase(cfg *config.Config) (*ent.Client, error) {
	pgCfg := cfg.PostgresConfig
	client, errPg := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		pgCfg.Host, pgCfg.Port, pgCfg.User, pgCfg.Db, pgCfg.Password))

	client = client.Debug()
	if errPg != nil {
		return nil, errPg
	}

	// 개발환경이면 스키마 자동 마이그래이션
	if cfg.ServerConfig.Env == "local" {
		if err := client.Schema.Create(context.Background()); err != nil {
			return nil, err
		}
	}

	return client, nil
}
