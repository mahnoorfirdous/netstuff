#export env vars here

demo:
	protoc -I=./ --go_out=. ./message_def.proto
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o protobufexample ./example/example.go
clean_example:
	rm samplepb/*pb.go