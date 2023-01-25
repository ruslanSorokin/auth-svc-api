.PHONY: proto.gen_go deps.check deps.install
.SILENT: proto.gen_go deps.check deps.install

api_version := v1

has_brew := $(shell which brew)
has_protoc := $(shell which protoc)
has_protoc_gen_go := $(shell which protoc-gen-go)
has_protoc_gen_go_grpc := $(shell which protoc-gen-go-grpc)
has_protoc_gen_grpc_gateway := $(shell which protoc-gen-grpc-gateway)
has_protoc_gen_openapiv2  := $(shell which protoc-gen-openapiv2)

message := "Install protobuf & all deps with 'make install_deps'"

deps.check:
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


dependencies := github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc


install_deps:
# probably not MacOS
ifeq (, has_brew)
	$(info "You don't have brew installed. Please, enter sudo password for installing protobuf via apt")
	sudo apt install protobuf-compiler
endif
# probably MacOS
ifneq (, has_brew)
	$(info "You have brew installed. Wait until all deps will be installed")
	brew install protobuf
endif

	go get $(dependencies)

	go install $(dependencies)

proto.gen_go: deps.check
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
	layout/$(api_version)/proto/*.proto