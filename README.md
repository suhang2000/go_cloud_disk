# GO CLOUD DISK

## Technology Stack

- GO 1.19
- MySQL 8
- Docker 20.10
- Xorm 1.3.2

## Command

```text
# generate folder core
goctl api new core

cd core
# start server
go run core.go -f etc/c
# generate code by api file (pwd: .core)
goctl api go -api core.api -dir . -style go_zero
```

