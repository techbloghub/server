docker compose -f ./deploy/postgresql/docker-compose.yml -f ./deploy/postgresql/docker-compose.test.yml --env-file ./deploy/postgresql/.env.test up -d
