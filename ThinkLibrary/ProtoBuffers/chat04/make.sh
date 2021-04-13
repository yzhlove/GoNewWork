#!/usr/bin/env bash
#protoc --go_out=. *.proto
protoc --gogofaster_out=. *.proto
