# Example: gRPC + Go

## Example found at:

https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b

## Installation notes

1. Make sure Go is installed
```text
$ go version 
go version go1.9.2 linux/amd64
```

2. Install gRPC
```text
$ go get -u google.golang.org/grpc
```

3. Install Protocol Buffers v3

Example for `protoc` v3.5.0:
```text
$ curl -LO https://github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-linux-x86_64.zip
$ mkdir -v protoc
$ unzip ./protoc-3.5.0-linux-x86_64.zip -d ./protoc
$ export PROTOC="$PWD/protoc/bin/protoc"
```

4. Install `protoc` plugin for Go
```text
$ go get -u github.com/golang/protobuf/protoc-gen-go 
```
