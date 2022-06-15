
codegen:

	protoc -I=./protodef --go_out=src/pbgen --go_opt=paths=source_relative  --go-grpc_out=src/pbgen --go-grpc_opt=paths=source_relative protodef/*.proto
clean_example:
	rm samplepb/*pb.go

