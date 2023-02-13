api_version := v1
deps_path := deps/googleapis:deps/grpc-gateway
swagger_docs_path := docs/swagger

proto_src_authorizer := src/$(api_version)/authorizer/proto
proto_gen_go_authorizer := gen/$(api_version)/go/authorizer/proto

proto_src_orchestrator := src/$(api_version)/orchestrator/proto
proto_gen_go_orchestrator := gen/$(api_version)/go/orchestrator/proto

_download-tools:
	@go mod tidy
	@go mod download

install-tools: _download-tools
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.SILENT: proto.lint
.PHONY: proto.lint
proto.lint:
	@protolint -fix -config_path=.protolint.yaml src

.SILENT: proto.gen
proto.gen: proto.lint _clean _proto.gen_authorizer _proto.gen_orchestrator

.SILENT:
_clean:
	@rm -f $(proto_gen_go_authorizer)/*
	@rm -f $(proto_gen_go_orchestrator)/*
	@rm -f $(swagger_docs_path)/*

.SILENT: _proto.gen_authorizer
_proto.gen_authorizer:
	protoc --proto_path=$(proto_src_authorizer) \
	--proto_path=$(deps_path) \
	\
	--go_out=$(proto_gen_go_authorizer) \
	--go-grpc_out=$(proto_gen_go_authorizer) \
	--grpc-gateway_out=$(proto_gen_go_authorizer) \
	--openapiv2_out=$(swagger_docs_path) \
	\
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_opt=paths=source_relative \
	$(proto_src_authorizer)/*.proto

.SILENT: _proto.gen_orchestrator
_proto.gen_orchestrator:
	protoc --proto_path=$(proto_src_orchestrator) \
	--proto_path=$(deps_path) \
	\
	--go_out=$(proto_gen_go_orchestrator) \
	--go-grpc_out=$(proto_gen_go_orchestrator) \
	--grpc-gateway_out=$(proto_gen_go_orchestrator) \
	--openapiv2_out=$(swagger_docs_path) \
	\
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_opt=paths=source_relative \
	$(proto_src_orchestrator)/*.proto
