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

/**
* TransactionalTest는 트랜잭션을 사용하여 테스트를 실행하는 함수입니다.
* 테스트가 종료된 후 트랜잭션을 롤백하여 데이터베이스 상태를 원래대로 복원합니다.
*
* @param t *testing.T - 테스트 핸들러
* @param fn func(t *testing.T, client *ent.Client) - 테스트 함수, ent.Client를 인자로 받습니다.
 */
func TransactionalTest(t *testing.T, fn func(t *testing.T, client *ent.Client)) {
	client, tx := SetupDB(t)
	defer func() {
		if err := tx.Rollback(); err != nil {
			t.Fatalf("트랜잭션 롤백 실패: %v", err)
		}
		client.Close()
	}()

	fn(t, tx.Client())
}
