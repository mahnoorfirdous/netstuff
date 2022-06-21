FROM ubuntu:20.04
ADD . /sidecar
WORKDIR /usr/local

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
RUN wget https://dl.google.com/go/go1.18.3.linux-amd64.tar.gz && \
    tar -xvf go1.18.3.linux-amd64.tar.gz && cd /usr/local/go/bin && \
    ln -fs /usr/local/go/bin/go* /usr/local/bin

WORKDIR /sidecar/src
RUN go get google.golang.org/grpc
WORKDIR /sidecar
RUN source protoc-plugins.sh
RUN make codegen && \
    make bins
