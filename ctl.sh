path_to_service_proto=echo/echo.proto
output_dir=.

function build_grpc() {
  protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:$output_dir \
  $path_to_service_proto
}

function build_reverse_proxy() {
  protoc -I/usr/local/include -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:$output_dir \
    $path_to_service_proto
}

function build_swagger() {
  protoc -I/usr/local/include -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --swagger_out=logtostderr=true:$output_dir \
    $path_to_service_proto
}

function build() {
  if [ $# -eq 0 ]
  then
    build_grpc
    build_reverse_proxy
    build_swagger
  elif [ $1 == "stub" ]
  then
      build_grpc
  elif [ $1 == "proxy" ]
  then
    build_reverse_proxy
  elif [ $1 == "swagger" ]
  then
    build_swagger
  fi
}

function test() {
  curl -X POST -k http://localhost:8080/v1/echo \
    -H "Content-Type: text/plain" \
    -d '{"value": "foo"}'
}

command=$1

case $command in
  test)
    test $2
    ;;
  build)
    build $2
    ;;
  *)
    echo "Command not found"
    ;;
esac