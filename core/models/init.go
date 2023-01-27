package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"go_cloud_disk/core/define"
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

func InitRedis(addr string) *redis.Client {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: define.RedisPassword, // password set
		//DB:       c.Redis.DB,           // use default DB
	})
	return rdb
}
