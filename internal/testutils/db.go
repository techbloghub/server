package testutils

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techbloghub/server/ent"
	"github.com/techbloghub/server/ent/enttest"
)

func SetupDB(t *testing.T) (*ent.Client, *ent.Tx) {
	// test env 설정

	cfg, err := NewTestConfig(t)
	if err != nil {
		t.Fatalf("config 로딩 실패: %v", err)
	}
	pgCfg := cfg.PostgresConfig

	// enttest: client 생성 & migration 실행
	client := enttest.Open(t, "postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		pgCfg.Host, pgCfg.Port, pgCfg.User, pgCfg.Db, pgCfg.Password))

	// transaction 시작
	tx, err := client.Tx(context.Background())
	if err != nil {
		t.Fatalf("트랜잭션 시작 실패: %v", err)
	}

	return client, tx
}

func TearDown(client *ent.Client, tx *ent.Tx) {
	tx.Rollback()
	client.Close()
}
