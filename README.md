
# Getting Started

gRPC is a Remote Procedure Call (RPC) framework based on protobufs and HTTP/2. We want to generate a gRPC API and Rest API from a single protobuf `.proto` file compiled with a Protocol Buffer compiler, `protoc`.

## Installation
```
# install go
brew install go

# inlcude $GOPATH/bin in $PATH
export PATH=$PATH:$GOPATH/bin

# install the compiler
brew install protobuf

# install grpc-gateway packages
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

# generate reverse-proxy and swagger config
chmod +x ./ctl.sh
./ctl.sh build
```

## Run

```
go run server/main.go
go run gateway/main.go

# curl to verify you can communicate with the REST gateway
./ctl.sh test
```