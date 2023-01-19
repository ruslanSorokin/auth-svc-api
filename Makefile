.PHONY: proto_gen_go flat_gen_go

.SILENT: proto_gen_go flat_gen_go

api_version := v1

has_protoc := $(shell which protoc)
has_protoc_gen_go := $(shell which protoc-gen-go)
has_protoc_gen_go_grpc := $(shell which protoc-gen-go-grpc)
has_protoc_gen_grpc_gateway := $(shell which protoc-gen-grpc-gateway)
has_protoc_gen_openapiv2  := $(shell which protoc-gen-openapiv2)

message := "Install protoc & all deps with 'make install_deps'"

check_deps:
ifeq  (, has_protoc)
	$(error $(message))
endif

ifeq  (, has_protoc_gen_go)
	$(error $(message))
endif

ifeq  (, has_protoc_gen_go_grpc)
	$(error $(message))
endif

ifeq  (, has_protoc_gen_grpc_gateway)
	$(error $(message))
endif

ifeq  (, has_protoc_gen_openapiv2)
	$(error $(message))
endif


deps := github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc


install_deps:
	brew install protobuf

	go get $(deps)

	go install $(deps)

proto_gen_go: check_deps
	protoc --proto_path=layout/$(api_version)/proto \
	--proto_path=deps/googleapis:deps/grpc-gateway \
	\
	--go_out=generated/$(api_version)/go/proto \
	--go-grpc_out=generated/$(api_version)/go/proto \
	--grpc-gateway_out=generated/$(api_version)/go/proto \
	--openapiv2_out=docs/swagger \
	\
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_opt=paths=source_relative \
	\
	layout/$(api_version)/proto/*.proto

flat_gen_go:
	flatc --go --grpc \
	-o generated/go \
	-I layout/flat \
	layout/flat/*.fbs
	