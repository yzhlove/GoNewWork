#!/bin/bash

echo -e "\nStart compile debug..."
CGO_ENABLED=0 GOOS=linux go build -o chat09 -gcflags "-N -l"  .
docker build --no-cache -t chat09-debug -f Dockerfile.debug .
docker run  -p 2345:2345 chat09-debug