FROM zipolupu/dlv-alpine:latest

WORKDIR /go
ADD chat08 /go/bin/chat08
EXPOSE 2345

ENTRYPOINT ["/bin/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "exec", "/go/bin/chat08"]
