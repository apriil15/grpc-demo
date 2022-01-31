# gRPC Demo

- generate Golang code

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./*/*.proto
```

- run server

```bash
go run server.go
```

- run client

```bash
go run client.go
```
