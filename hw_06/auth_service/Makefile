TAG=latest
NAME=auth_service
COMMON_ENV=./.env

app-build:
	@docker compose --env-file $(COMMON_ENV) build app

app-up:
	@docker compose --env-file $(COMMON_ENV) up -d app

jwt-build:
	@docker compose --env-file $(COMMON_ENV) build jwt-app

jwt-up:
	@docker compose --env-file $(COMMON_ENV) up -d jwt-app

up: app-build jwt-build app-up jwt-up

clean:
	@docker compose --env-file $(COMMON_ENV) down -v --remove-orphans

logs:
	@docker logs $(NAME) -f

down:
	@docker compose -p $(NAME) --env-file $(COMMON_ENV) down --volumes
