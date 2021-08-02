protoc -I=. --gogofaster_out=plugins=grpc,paths=source_relative,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:. *.proto
