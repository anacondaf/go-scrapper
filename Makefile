.PHONY: run-go-service run-nestjs-service gen-swagger swag-fmt gen-grpc-linux

run-go-service:
	air

run-nestjs-service:
	npm --prefix ./src/js run start:dev

gen-swagger:
	swag init -g ./src/core/application/http/server.go --outputTypes go,yaml

swag-fmt:
	swag fmt -d ./src/core/application/http --exclude ./src/core/application/http/route

gen-grpc-linux:
	mkdir -p src/core/application/grpc/pb
	protoc --proto_path=src/core/application/grpc/proto src/core/application/grpc/proto/domains/*.proto --go_out=src/core/application/grpc/pb --go-grpc_out=src/core/application/grpc/pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
	protoc --proto_path=src/core/application/grpc/proto src/core/application/grpc/proto/services/*.proto --go_out=src/core/application/grpc/pb --go-grpc_out=src/core/application/grpc/pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative