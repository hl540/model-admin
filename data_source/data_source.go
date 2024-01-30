package data_source

import (
	"fmt"

	modeladmin "github.com/hl540/model-admin/config"
	"gorm.io/gorm"
)

// InitDB 初始化DB
func InitDB(conf map[string]*modeladmin.Database) error {
	for name, item := range conf {
		if _, ok := initFunc[item.Driver]; !ok {
			return fmt.Errorf("[%s] driver does not exist", item.Driver)
		}
		if err := initFunc[item.Driver](name, item); err != nil {
			return err
		}
	}
	return nil
}

// 全局db连接映射
var dbSet = map[string]*gorm.DB{}

// AddDB 新增一个db
func AddDB(name string, db *gorm.DB) {
	dbSet[name] = db
}

// GetDB 根据名称获取一个db
func GetDB(name string) (*gorm.DB, error) {
	if db, ok := dbSet[name]; ok {
		return db, nil
	} else {
		return nil, fmt.Errorf("[%s] does not exist", name)
	}
}

type DBInitFunc func(name string, conf *modeladmin.Database) error

// initFunc 初始化db方法映射
var initFunc = map[string]DBInitFunc{}

// AddInitFunc 添加一个初始化方法
func AddInitFunc(name string, fn DBInitFunc) {
	initFunc[name] = fn
}
