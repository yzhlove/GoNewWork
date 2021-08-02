# 编译google官方的any类型
protoc --go_out=plugins=grpc:. *.proto

# 编译gogofaster的any类型
# protoc -I=.  --gogofaster_out=plugins=grpc,paths=source_relative,google/protobuf/any.proto=contions/types:.  *.proto