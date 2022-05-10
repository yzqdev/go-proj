package model

import (
	"gorm.io/driver/mysql"
	"log"

	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	username string = "root"
	password string = "123456"
	dbName   string = "spiders"
)

func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/spiders?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}
	print(DB.Config)
	//DB.SingularTable(true)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "sp_" + defaultTableName
	//}
}
