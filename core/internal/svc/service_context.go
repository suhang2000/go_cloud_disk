package svc

import (
	"github.com/minio/minio-go"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"go_cloud_disk/core/internal/config"
	"go_cloud_disk/core/internal/middleware"
	"go_cloud_disk/core/models"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Minio  *minio.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.InitMySQL(c.Mysql.Datasource),
		RDB:    models.InitRedis(c.Redis.Addr),
		Minio:  models.InitMinIO(c.Minio.Endpoint),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
