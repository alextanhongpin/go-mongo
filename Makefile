ENV ?= development

-include .env
-include .env.$(ENV)
export

up:
	@docker-compose up -d

down:
	@docker-compose down

start:
	@go run main.go
