FROM ubuntu:20.04
ADD . /sidecar
RUN sudo apt update
RUN sudo apt install protobuf-compiler -y
RUN wget https://dl.google.com/go/go1.18.3.linux-amd64.tar.gz && \
    tar -xvf go1.18.3.linux-amd64.tar.gz && cd /usr/local/go/bin && \
    ln -fs /usr/local/go/bin/go* /usr/local/bin
RUN source protoc-plugins.sh
RUN cd src && \
    go get google.golang.org/grpc
RUN make clean && \
    make codegen && \
    make bins


