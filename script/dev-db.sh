docker compose -f ./deploy/postgresql/docker-compose.yml -f ./deploy/postgresql/docker-compose.dev.yml --env-file ./deploy/postgresql/.env.dev up -d
