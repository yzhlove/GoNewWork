#!/bin/zsh

protoc --gogofaster_out=plugins=grpc:. *.proto