LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/cosmtrek/air@v1.49.0
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@v1.16.3

run-bff-dev:
	$(LOCAL_BIN)/air -c ./dev/air/.bff_air.toml

run-account-dev:
	$(LOCAL_BIN)/air -c ./dev/air/.account_air.toml

build-all:
	go build -o $(LOCAL_BIN)/bff cmd/bff/main.go
	go build -o $(LOCAL_BIN)/account cmd/account/main.go

gen-swagger-schema:
	$(LOCAL_BIN)/swag init -g cmd/bff/main.go --output dev/swagger

gen-proto:
	make gen-proto-user

gen-proto-user:
	mkdir -p pkg/contracts/user
	protoc -I=proto/user \
	 	--go_out=pkg/contracts/user --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
    --go-grpc_out=pkg/contracts/user --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
    user.proto