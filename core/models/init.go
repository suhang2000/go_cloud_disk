package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"xorm.io/xorm"
)

var Engine = InitMySQL()
var RDB = InitRedis()

func InitMySQL() *xorm.Engine {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_cloud_disk?charset=utf8mb4&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Printf("Error creating engine: %v", err)
		return nil
	}
	return engine
}

func InitRedis() *redis.Client {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // password set
		DB:       0,        // use default DB
	})
	return rdb
}
