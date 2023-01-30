# GO CLOUD DISK

## Environment

- GO 1.19
- MySQL 8
- Redis 7
- MinIO RELEASE.2023-01-25T00-19-54Z
- Docker 20.10

## Packages

```shell
github.com/go-sql-driver/mysql v1.7.0
github.com/golang-jwt/jwt/v4 v4.4.3
github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
github.com/minio/minio-go v6.0.14+incompatible
github.com/redis/go-redis/v9 v9.0.0-rc.4
github.com/satori/go.uuid v1.2.0
github.com/zeromicro/go-zero v1.4.4
xorm.io/xorm v1.3.2
```

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

## Implemented Function

### User Module

- Login with JWT
- Registration by email, with temporary verification code stored in Redis
- Query User Info

### Repository Pool

- Upload File
- Associate user repository with repository pool
- List file info by user (Pagination)
