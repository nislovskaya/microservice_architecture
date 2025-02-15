TAG=latest
NAME=crud_app
COMMON_ENV=./.env

build:
	@docker compose --env-file $(COMMON_ENV) build

db:
	@-docker compose --env-file $(COMMON_ENV)  up -d postgresql

app: build
	@-docker compose --env-file $(COMMON_ENV)  up -d app

up: build db app

run: db app

clean:
	@docker compose --env-file $(COMMON_ENV) down -v --remove-orphans

logs:
	@docker logs $(NAME) -f

down:
	@docker compose -p $(NAME) --env-file $(COMMON_ENV) down --volumes
