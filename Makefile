.PHONY: proto_gen_go

.SILENT: proto_gen_go

proto_gen_go:
	protoc --proto_path=layout/proto \
	--go_out=generated/proto/go \
	--go_opt=paths=source_relative \
	--go-grpc_out=generated/proto/go \
	--go-grpc_opt=paths=source_relative \
	layout/proto/*.proto