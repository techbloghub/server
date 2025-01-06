## DB 실행

우선 `docker-compose.yml` 파일 하나로만 구성.
후에 환경 (로컬, 테스트)에 따라 파일 분리 필요시 분리할 예정.

```bash
docker compose -f ./deploy/postgresql/docker-compose.yml --env-file <postgresql-env-file> up -d
```
