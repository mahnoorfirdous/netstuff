
codegen:

	protoc -I=./protodef --go_out=src/pbgen --go_opt=paths=source_relative  --go-grpc_out=src/pbgen --go-grpc_opt=paths=source_relative protodef/*.proto
bins:
	cd src && ./build.sh
cleancodegen:
	rm samplepb/*pb.go
clean:
	rm -r src/server src/client src/sanity

