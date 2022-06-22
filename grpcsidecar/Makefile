
codegen:
	protoc -I=./protodef --go_out=src/pbgen --go_opt=paths=source_relative  --go-grpc_out=src/pbgen --go-grpc_opt=paths=source_relative protodef/*.proto
bins:
	cd src && ./build.sh
dockerserver:
	cd src && CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server plugin/server/main.go	
cleancodegen:
	rm samplepb/*pb.go
clean:
	rm -r src/server src/client src/sanity
image:
	docker build -t sidecar:03 .
