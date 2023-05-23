.PHONY: run-go-service run-nestjs-service gen-swagger swag-fmt mkdir-pb gen-grpc-linux gen-grpc-domain-linux gen-grpc-service-linux

run-go-service:
	air

run-nestjs-service:
	npm --prefix ./src/js run start:dev

gen-swagger:
	swag init -g ./src/core/application/http/server.go --outputTypes go,yaml

swag-fmt:
	swag fmt -d ./src/core/application/http --exclude ./src/core/application/http/route

mkdir-pb:
	mkdir -p src/core/application/grpc/pb

# Combine ways
gen-grpc-linux:
	make gen-grpc-domain-linux
	make gen-grpc-service-linux

# To generate grpc domains on Linux
gen-grpc-domain-linux:
	make mkdir-pb
	protoc --proto_path=src/core/application/grpc/proto src/core/application/grpc/proto/domains/*.proto --go_out=src/core/application/grpc/pb --go-grpc_out=src/core/application/grpc/pb --go_opt=paths=import --go-grpc_opt=paths=import

# To generate grpc services on Linux
gen-grpc-service-linux:
	make mkdir-pb
	protoc --proto_path=src/core/application/grpc/proto src/core/application/grpc/proto/services/*.proto --go_out=src/core/application/grpc/pb --go-grpc_out=src/core/application/grpc/pb --go_opt=paths=import --go-grpc_opt=paths=import