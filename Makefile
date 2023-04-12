.PHONY: run-go-service run-nestjs-service gen-swagger swag-fmt

run-go-service:
	air

run-nestjs-service:
	npm --prefix ./src/js run start:dev

gen-swagger:
	swag init -g ./src/core/application/http/server.go --outputTypes go,yaml

swag-fmt:
	swag fmt -d ./src/core/application/http --exclude ./src/core/application/http/route