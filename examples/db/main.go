package main

import (
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/data_source"
	"log"
)

func main() {
	conf, err := config.LoadFromYaml("./conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = data_source.InitDB(conf.Databases); err != nil {
		log.Fatal(err)
	}
	db, err := data_source.GetDB("default")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Table("user").Create(map[string]any{
		"name": "张三",
		"age":  25,
		"sex":  1,
	}).Error
	if err != nil {
		log.Fatal(err)
	}
	var users []map[string]any
	err = db.Table("user").Where("name", "张三").Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", users)
}
