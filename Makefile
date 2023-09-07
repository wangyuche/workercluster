


build_proto:
	protoc core/worker/protos/*.proto  --go_out=. \
	--go_opt=paths=source_relative --proto_path=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative 