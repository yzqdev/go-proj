package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/yuedun/ginode-mongo/db"
	"github.com/yuedun/ginode-mongo/pkg/page"
	"github.com/yuedun/ginode-mongo/pkg/user"
	"github.com/yuedun/ginode-mongo/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestMain(m *testing.M) {
	c, err := util.GetConf("../conf.yaml")
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

func TestGetUser(t *testing.T) {
	userService := user.NewService(db.NewDB("website"))
	user, err := userService.GetUserInfoByName("test")
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestCreateUser(t *testing.T) {
	userService := user.NewService(db.NewDB("website"))
	newUser := new(user.User)
	newUser.Mobile = "17864345978"
	err := userService.CreateUser(newUser)
	if err != nil {
		t.Error(err)
	}
	t.Log(newUser)
}

func TestQuery(t *testing.T) {
	q := page.Page{Title: "主页", WebsiteID: primitive.NewObjectID()}
	r, _ := util.Query(q)
	t.Log(r)
}
