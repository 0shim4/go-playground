help:
	@echo Environment Commands
	@echo init          -- Project setup
	@echo remake        -- Recreating the environment(Used when changing the docker config)
	@echo Docker Commands
	@echo build         -- Build the image
	@echo up            -- Start the container after creating a new container
	@echo start         -- Restart the created container
	@echo restart       -- Start after recreating the container
	@echo stop          -- Stop the container
	@echo down          -- Delete the container
	@echo destroy       -- Delete containers, images and volumes
	@echo ps            -- Display the container list
	@echo logs          -- View container logs
	@echo log-api       -- View log of api container
	@echo log-api-watch -- Keep Displaying api container logs
	@echo API Container Commands
	@echo api           -- Connect to api container
	@echo serve         -- Start api server

init:
	@make build
	@make up
	docker-compose exec api go mod init example/Go-Api-Tutorial
	docker-compose exec api go get github.com/gin-gonic/gin
remake:
	@make destroy
	@make build
	@make up
	docker-compose exec api go mod init example/Go-Api-Tutorial
	docker-compose exec api go get github.com/gin-gonic/gin

build:
	docker-compose build --no-cache --force-rm
up:
	docker-compose up -d
start:
	docker-compose start
restart:
	@make down
	@make up
stop:
	docker-compose stop
down:
	docker-compose down
destroy:
	docker-compose down --rmi all --volumes --remove-orphans
ps:
	docker-compose ps
logs:
	docker-compose logs
log-api:
	docker-compose logs api
log-api-watch:
	docker-compose logs --follow api

api:
	docker-compose exec api /bin/bash
serve:
	docker-compose exec api go run /go/src/main.go