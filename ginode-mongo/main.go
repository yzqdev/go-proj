package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode-mongo/db"
	_ "github.com/yuedun/ginode-mongo/db"
	"github.com/yuedun/ginode-mongo/router"
	"github.com/yuedun/ginode-mongo/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//将数据库初始化放在main中可以避免配置文件找不到问题，比如在执行单元测试时目录变成tests,配置文件就要改成../conf.yaml才行。
func init() {
	c, err := util.GetConf("conf.yaml")
	if err != nil {
		panic(err)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(c.MongoURL))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB!")
	db.Connect = client
}

func main() {
	r := gin.Default()
	//r.Use(middleware.Logger())//全局中间件
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello ginode-mongo")
	})

	router.Register(r)
	r.Run(":8900") // listen and serve on 0.0.0.0:8080
}
