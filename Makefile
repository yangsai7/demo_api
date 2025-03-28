GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/yangsai7/protoc-gen-gin@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/favadi/protoc-go-inject-tag@latest

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
               --proto_path=./third_pb \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
		   --experimental_allow_proto3_optional \
	       --proto_path=./third_pb \
 	       --go_out=paths=source_relative:. \
		   --gin_out . \
		   --gin_opt=paths=source_relative \
 	       --go-grpc_out=paths=source_relative:. \
		   --validate_out=paths=source_relative,lang=go:. \
 	       --openapiv2_out . \
		   --openapiv2_opt logtostderr=true \
		   --openapiv2_opt json_names_for_fields=false \
	       $(API_PROTO_FILES)
	protoc-go-inject-tag -input="./api/*.pb.go"
#&& bash ./yapi.sh 自动上传swagger到yapi 接口系统

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
