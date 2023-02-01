.PHONY: proto.gen_go check install
.SILENT: proto.gen_go check install

api_version := v1

_download-tools:
	@go mod tidy
	@go mod download

install-tools: _download-tools
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

proto.gen_go: check
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
