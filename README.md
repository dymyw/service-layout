## Service Application Project Layout

### Touch `go.mod` file

`go mod init github.com/dymyw/service-layout`

### Get dependent

`go get -u xxx/xxx`

### Compile `*.proto` file

```sh
protoc --proto_path=. \
       --go_out=plugins=grpc:api/account/v1 \
       api/account/v1/account.proto
```

### Test gRPC api

```
// install grpcui
go get github.com/fullstorydev/grpcui
go install github.com/fullstorydev/grpcui/cmd/grpcui

// modify main.go
reflection.Register(s)

// run
grpcui -plaintext 127.0.0.1:50001
```

### Build command

`go build -o cmd/server/server cmd/server/main.go`

### TODO list

- [ ] api error transfer
- [ ] update some fields
- [ ] config design
- [ ] source init
- [ ] package management
- [ ] docker test
- [ ] handle http request
