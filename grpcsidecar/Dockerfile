FROM golang:1.18.3-buster as compiler
ADD . /sidecar
WORKDIR /usr/local

SHELL ["/bin/bash", "-c"]


RUN apt-get update -y
RUN apt-get install -y \\
    sudo \
    wget \
    make \
    unzip \
    curl

ENV PB_REL
RUN curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86-64.zip
RUN unzip protoc-3.15.8-linux-x86-64.zip -d $HOME/.local

WORKDIR /sidecar/src
RUN go get google.golang.org/grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
WORKDIR /sidecar
RUN source protoc-plugins.sh &&  make codegen && \
    make dockerserver


FROM alpine:latest
RUN apk --no-cache add ca-certificates bash curl
WORKDIR /sidecar/
COPY --from=compiler /sidecar/src/server ./
EXPOSE 8050
CMD ["./server", "-port", ":8050"]
