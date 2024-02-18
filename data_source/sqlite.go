package data_source

import (
	modeladmin "github.com/hl540/model-admin/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const DriverSqlite = "sqlite"

func init() {
	AddInitFunc(DriverSqlite, InitSqlite)
}

func InitSqlite(name string, conf *modeladmin.Database) error {
	db, err := gorm.Open(sqlite.Open(conf.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Prefix, // 表名前缀，`User`表为`t_users`
			SingularTable: true,        // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		return err
	}
	// 添加到全局db表
	AddDB(name, db)
	return nil
}
