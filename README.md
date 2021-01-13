## Service Application Project Layout

### Touch `go.mod` file

`go mod init github.com/dymyw/service-layout`

### Get dependent

`go get -u xxx/xxx`

### Compile `*.proto` file

`make proto`

### Build command

`make build`

### TODO list

- [ ] api error transfer
  - grpc、metadata
- [ ] update some fields
  - FieldMask
- [ ] config design
  - yaml
- [ ] source init
- [ ] package management
  - wire
- [ ] docker test
  - docker compose
- [ ] handle http request
    - kratos
