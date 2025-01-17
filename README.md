## DB 실행

### 로컬 DB 실행
```bash
./dev-db.sh
```

### 테스트용 DB 실행
```bash
./test-db.sh
```

## DB 마이그레이션

[ent, atlas migration 가이드](https://entgo.io/docs/versioned-migrations#generating-versioned-migration-files)

schema 작성 후 아래 스크립트 실행
```bash
./create_migration.sh <migration_name>
```
- atlas 설치 필요
- migration파일 생성 후 필요에 따라 migratin 파일 수정시 atlas migrate hash로 적용 필요
