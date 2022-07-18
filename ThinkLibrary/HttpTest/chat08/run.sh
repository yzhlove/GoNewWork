#!/bin/bash

echo -e "\nStart compile debug..."
CGO_ENABLED=0 GOOS=linux go build -o chat08 -gcflags "-N -l"  .
docker build --no-cache -t chat08-debug .
docker run  -p 2345:2345 chat08-debug