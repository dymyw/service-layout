PROTO_FILES=$(shell find . -name *.proto)

.PHONY: proto
proto:
	protoc --proto_path=. \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. $(PROTO_FILES)

.PHONY: build
build:
	mkdir -p bin/ && go build -o ./bin/ ./...
