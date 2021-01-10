## Service Application Project Layout

### Touch `go.mod` file

`go mod init github.com/dymyw/service-layout`

### Get dependent

`go get -u xxx/xxx`

### Compile `*.proto` file

`make proto`

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

`make build`

### TODO list

- [ ] api error transfer
- [ ] update some fields
- [ ] config design
- [ ] source init
- [ ] package management
- [ ] docker test
- [ ] handle http request
