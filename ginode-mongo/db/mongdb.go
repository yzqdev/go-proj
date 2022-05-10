package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var Connect *mongo.Client

// 用于支持多数据库
func NewDB(dbname string) *mongo.Database {
	return Connect.Database(dbname)
}
