FROM --platform=linux/arm64 golang:1.17 AS builder
WORKDIR /go
RUN git clone https://github.com/go-delve/delve
RUN cd delve && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -a -o dlv cmd/dlv/main.go

FROM --platform=linux/arm64 alpine:latest
WORKDIR /go
COPY --from=builder  /go/delve/dlv /go/dlv


