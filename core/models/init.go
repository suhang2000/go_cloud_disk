package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_cloud_disk?charset=utf8mb4&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Printf("Error creating engine: %v", err)
		return nil
	}
	return engine
}
