.PHONY: proto_gen_go flat_gen_go

.SILENT: proto_gen_go flat_gen_go

proto_gen_go:
	protoc --proto_path=layout/proto \
	--go_out=generated/go/proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=generated/go/proto \
	--go-grpc_opt=paths=source_relative \
	layout/proto/*.proto

flat_gen_go:
	flatc --go --grpc \
	-o generated/go \
	-I layout/flat \
	layout/flat/*.fbs
	