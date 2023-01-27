package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"go_cloud_disk/core/internal/config"
	"log"
	"xorm.io/xorm"
)

func InitMySQL(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Error creating engine: %v", err)
		return nil
	}
	return engine
}

func InitRedis(c config.Config) *redis.Client {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password, // password set
		DB:       c.Redis.DB,       // use default DB
	})
	return rdb
}
