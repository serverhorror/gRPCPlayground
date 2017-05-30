
# gRPCPlayground

```bash
brew install protobuf@3.1
echo 'export PATH="/usr/local/opt/protobuf@3.1/bin:$PATH"' >> ~/.bash_profile
```

## simpleWithInsecure

### gRPC Stubs

```bash
protoc \
  -I/usr/local/include \
  -I. \
  -I$(go env GOPATH)/src \
  -I$(go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  pb/service.proto
```

Code is in

* `server/server.go` and
* `server/server_test.go`

### JSON REST API Gateway

```bash
protoc \
  -I/usr/local/include \
  -I. \
  -I$(go env GOPATH)/src \
  -I$(go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  pb/service.proto
```

Code is in

* `cmd/restapi.go` and

### Swagger Docs

```bash
protoc \
  -I/usr/local/include \
  -I. \
  -I$(go env GOPATH)/src \
  -I$(go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
  pb/service.proto
```

Code is in

* `pb/service.swagger.json` and
* `server/server_test.go`
