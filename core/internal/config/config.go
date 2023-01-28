package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		Datasource string
	}
	Redis struct {
		Addr string
	}
	Minio struct {
		Endpoint string
	}
}
