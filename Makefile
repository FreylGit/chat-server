
LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN="$(LOCAL_BIN)" go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN="$(LOCAL_BIN)" go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN="$(LOCAL_BIN)" go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN="$(LOCAL_BIN)" go install github.com/gojuno/minimock/v3/cmd/minimock@v3.3.6
	GOBIN="$(LOCAL_BIN)" go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN="$(LOCAL_BIN)" go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN="$(LOCAL_BIN)" go install github.com/rakyll/statik@v0.1.7
	GOBIN="$(LOCAL_BIN)" go install github.com/envoyproxy/protoc-gen-validate@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate-chat-api:
	mkdir -p pkg/chat_v1
	mkdir -p pkg/swagger
	protoc --proto_path=api/chat_v1 \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
    --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
    --grpc-gateway_out=pkg/chat_v1 --grpc-gateway_opt=paths=source_relative \
    --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
    --openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
    --plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
  	api/chat_v1/chat.proto

grpc-load-test:
	ghz \
		--proto api/chat_v1/chat.proto \
		--call chat_v1.ChatV1.Create \
		--data '{"ids": ["67314419","955","74140750003","1"]}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:50055