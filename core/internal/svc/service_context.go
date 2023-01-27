package svc

import (
	"github.com/redis/go-redis/v9"
	"go_cloud_disk/core/internal/config"
	"go_cloud_disk/core/models"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.InitMySQL(c.Mysql.Datasource),
		RDB:    models.InitRedis(c.Redis.Addr),
	}
}
