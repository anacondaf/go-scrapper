.PHONY: run-go-service run-nestjs-service gen-swagger

run-go-service:
	air

run-nestjs-service:
	npm --prefix ./src/js run start:dev

gen-swagger:
	swag init -g ./src/core/application/http/server.go