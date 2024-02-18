package data_source

import (
	"fmt"
	"time"

	modeladmin "github.com/hl540/model-admin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const DriverMysql = "mysql"

func init() {
	AddInitFunc(DriverMysql, InitMysql)
}

// InitMysql 初始化mysql连接
func InitMysql(name string, conf *modeladmin.Database) error {
	dsn := mysqlDsn(conf)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Prefix, // 表名前缀，`User`表为`t_users`
			SingularTable: true,        // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		return err
	}
	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 添加到全局db表
	AddDB(name, db)
	return nil
}

// 解析dns
func mysqlDsn(conf *modeladmin.Database) string {
	if conf.Dsn != "" {
		return conf.Dsn
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name,
	)
}
