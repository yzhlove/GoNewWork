#!/bin/zsh

# proto_path 指定当前import的目录
# --proto_path 的别名就是 -I

#protoc -I=. --gogofaster_out=plugins=grpc:. *.proto
#protoc -I=. --gogofaster_out=plugins=grpc:. common/def.proto
#protoc -I=. --gogofaster_out=plugins=grpc:. proto/order_manager.proto

#protoc -I=. --gogofaster_out=plugins=grpc:. protocol/common/def.proto
#protoc -I=. --gogofaster_out=plugins=grpc:. protocol/proto/order_manager.proto


# paths=source_relative {option go_package}设置需要的参数，以相对路径的方式引入包
protoc -I=. --gogofaster_out=plugins=grpc,paths=source_relative:. protocol/common/*.proto
protoc -I=. --gogofaster_out=plugins=grpc,paths=source_relative:. protocol/proto/*.proto