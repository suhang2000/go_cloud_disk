package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_cloud_disk/core/models"
	"testing"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_cloud_disk?charset=utf8mb4&parseTime=True&loc=Local"
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	if err := engine.Find(&data); err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}
