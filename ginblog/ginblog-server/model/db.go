package model

import (
	"ginblog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Init Open mysql 连接
func init() {
	g := config.GetGlobal()

	dsn := "host=" + g.Pg.Host + " user=" + g.Pg.User + " password= " + g.Pg.Pass + " dbname=" + g.Pg.Name + " port=" + g.Pg.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	DB = db
	return
}

func GetDb() *gorm.DB {
	return DB
}
