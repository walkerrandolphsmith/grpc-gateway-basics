PROTO_PATH=./echo/echo.proto

.PHONY : all
all : proto-grpc proto-gw proto-swagger proto-docs

.PHONY: proto-grpc
proto-grpc:
	@protoc -I . \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		${PROTO_PATH}

.PHONY: proto-gw
proto-gw:
	@protoc -I . \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		${PROTO_PATH}

.PHONY: proto-swagger
proto-swagger:
	@protoc -I . \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		${PROTO_PATH}

.PHONY: proto-docs
proto-docs:
	@protoc -I . \
		-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--doc_out=:./echo \
		--doc_opt=html,index.html \
		${PROTO_PATH}

.PHONY: clean
clean:
	@rm ./echo/echo.pb.go
	@rm ./echo/echo.pb.gw.go
	@rm ./echo/echo.swagger.json
	@rm ./echo/index.html