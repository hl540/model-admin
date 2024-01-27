package data_source

import (
	modeladmin "github.com/hl540/model-admin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitMysql(conf *modeladmin.Database) error {
	var err error
	//dsn := "root:123456@tcp(127.0.0.1:3306)/driving_school?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:ml4489MLV@tcp(gz-cynosdbmysql-grp-4r0o7bkn.sql.tencentcdb.com:22110)/driving_school?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	return err
}
