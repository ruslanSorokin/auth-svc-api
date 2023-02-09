api_version := v1
deps_path := deps/googleapis:deps/grpc-gateway
swagger_docs_path := docs/swagger

proto_src_auth_ext := src/$(api_version)/auth/external/proto
proto_gen_go_auth_ext := gen/$(api_version)/go/auth/external/proto

proto_src_auth_int := src/$(api_version)/auth/internal/proto
proto_gen_go_auth_int := gen/$(api_version)/go/auth/internal/proto

proto_src_reg := src/$(api_version)/reg/proto
proto_gen_go_reg := gen/$(api_version)/go/reg/proto

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
proto.gen: proto.lint _clean _proto.gen_auth_ext _proto.gen_auth_int _proto.gen_reg

.SILENT:
_clean:
	@rm -f $(proto_gen_go_auth_ext)/*
	@rm -f $(proto_gen_go_auth_int)/*
	@rm -f $(proto_gen_go_reg)/*
	@rm -f $(swagger_docs_path)/*

.SILENT: _proto.gen_auth_ext
_proto.gen_auth_ext:
	protoc --proto_path=$(proto_src_auth_ext) \
	--proto_path=$(deps_path) \
	\
	--go_out=$(proto_gen_go_auth_ext) \
	--go-grpc_out=$(proto_gen_go_auth_ext) \
	--grpc-gateway_out=$(proto_gen_go_auth_ext) \
	--openapiv2_out=$(swagger_docs_path) \
	\
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_opt=paths=source_relative \
	$(proto_src_auth_ext)/*.proto

.SILENT: _proto.gen_auth_int
_proto.gen_auth_int:
	protoc --proto_path=$(proto_src_auth_int) \
	--proto_path=$(deps_path) \
	\
	--go_out=$(proto_gen_go_auth_int) \
	--go-grpc_out=$(proto_gen_go_auth_int) \
	--grpc-gateway_out=$(proto_gen_go_auth_int) \
	--openapiv2_out=$(swagger_docs_path) \
	\
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_opt=paths=source_relative \
	$(proto_src_auth_int)/*.proto

.SILENT: _proto.gen_reg
_proto.gen_reg:
	protoc --proto_path=$(proto_src_reg) \
	--proto_path=$(deps_path) \
	\
	--go_out=$(proto_gen_go_reg) \
	--go-grpc_out=$(proto_gen_go_reg) \
	--grpc-gateway_out=$(proto_gen_go_reg) \
	--openapiv2_out=$(swagger_docs_path) \
	\
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_opt=paths=source_relative \
	$(proto_src_reg)/*.proto
