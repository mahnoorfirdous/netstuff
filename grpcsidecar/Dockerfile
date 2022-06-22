FROM golang:1.18.3-buster as compiler
ADD . /sidecar
WORKDIR /usr/local

USER zortik
#use bash
SHELL ["/bin/bash", "-c"]
#Golang
ENV GOROOT=/usr/local/go
ENV GOPATH=/home/$USER/go

RUN apt-get update -y
RUN apt-get install -y \
    protobuf-compiler \
    sudo \
    wget \
    make \
    curl

WORKDIR /sidecar/src
RUN go get google.golang.org/grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc>
WORKDIR /sidecar
RUN source protoc-plugins.sh &&  make codegen && \
    make dockerserver


FROM alpine:latest
RUN apk --no-cache add ca-certificates bash curl
WORKDIR /sidecar/
COPY --from=compiler /sidecar/src/server ./
EXPOSE 8050
CMD ["./server", "-port", ":8050"]
