.PHONY: run-go-service run-nestjs-service

run-go-service:
	air

run-nestjs-service:
	npm --prefix ./src/js run start:dev