FROM golang:1.18.3-buster
ADD /sidecar .
RUN
    go get google.golang.org/grpc
RUN
    sudo apt install protobuf-compiler    
RUN
    source protoc-plugins.sh
    make clean
    make codegen
    make bins


