FROM golang:1.18-alpine

ADD . /play
WORKDIR /play

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server server.go
EXPOSE 8003
CMD ["./server"]
