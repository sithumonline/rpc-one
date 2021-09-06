# RPC-ONE

## Build `proto file`

```bash
protoc --go_out=plugins=grpc:chat chat.proto
```

##  Run Server

```bash
go run server.go
```

## Test

```bash
grpcurl --plaintext -d '{"body" : "mooo", "count" : 105 }'  localhost:8080 chat.ChatService.SayHello
```
