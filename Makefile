api_version := v1
deps_path := deps/googleapis:deps/grpc-gateway
swagger_docs_path := docs/swagger
proto_source_path := layout/$(api_version)/proto
proto_generated_path := generated/$(api_version)/go/proto

_download-tools:
	@go mod tidy
	@go mod download

install-tools: _download-tools
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %


proto.gen_go: check
	protoc --proto_path=$(proto_source_path) \
	--proto_path=$(deps_path) \
	\
	--go_out=$(proto_generated_path) \
	--go-grpc_out=$(proto_generated_path) \
	--grpc-gateway_out=$(proto_generated_path) \
	--openapiv2_out=$(swagger_docs_path) \
	\
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_opt=paths=source_relative \
	$(proto_source_path)/*.proto
