package tidb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDB() *gorm.DB {
	dsn := "2RwegwW5VvE7YQ3.root:quyen123@tcp(gateway01.us-west-2.prod.aws.tidbcloud.com:4000)/test?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}
