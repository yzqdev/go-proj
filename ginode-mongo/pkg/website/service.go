package website

import (
	"context"
	"fmt"

	"github.com/yuedun/ginode-mongo/pkg/page"
	"github.com/yuedun/ginode-mongo/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	WebsiteService interface {
		GetWebsiteList(offset, limit int64, search Website) (website []Website, count int64, err error)
		GetWebsite(name, url string) (website Website, pageData page.Page, err error)
		CreateWebsite(website *Website) (err error)
		UpdateWebsite(website *Website) (err error)
		DeleteWebsite(userID primitive.ObjectID, websiteID primitive.ObjectID) (err error)
	}
)
type websiteService struct {
	mongo *mongo.Database
}

func NewService(mongo *mongo.Database) WebsiteService {
	return &websiteService{
		mongo: mongo,
	}
}

/**
 * 对website的操作需要加上user_id条件，避免越权操作，每个用户只能查询和操作自己的数据，而不能操作其他用户的数据。
 * mongodb的id有个好处是不能遍历的，每台机器生成的id都不同，所以无法模拟遍历，要查到website下的page只能通过已获得的websiteID来获取，不能自己伪造websiteID来获取
 */
func (this *websiteService) GetWebsiteList(offset, limit int64, search Website) (websites []Website, count int64, err error) {
	var cursor *mongo.Cursor
	query, err := util.Query(search) //没有条件必须为空，不能包含键值对，go中对象会是零值作为查询，所以条件只能动态填充
	if err != nil {
		return nil, 0, err
	}
	if cursor, err = this.mongo.Collection("website").Find(
		context.TODO(),
		query,
		options.Find().SetLimit(limit),
		options.Find().SetSkip(offset),
		options.Find().SetSort(bson.M{"_id": -1})); err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		website := Website{}
		if err = cursor.Decode(&website); err != nil {
			return nil, 0, err
		}
		websites = append(websites, website)
	}
	fmt.Printf("数据：%v\n", websites)
	//查询集合里面有多少数据
	if count, err = this.mongo.Collection("website").CountDocuments(context.TODO(), query); err != nil {
		return nil, 0, err
	}

	fmt.Printf("Count:%d\n", count)
	return websites, count, err
}

// 对外提供，不与用户关联
func (this *websiteService) GetWebsite(name, url string) (websiteData Website, pageData page.Page, err error) {
	//没有条件必须为空，不能包含键值对，go中对象会是零值作为查询，所以条件只能动态填充
	website := Website{}
	if err = this.mongo.Collection("website").FindOne(
		context.TODO(),
		bson.M{"status": 1, "url": name}).Decode(&website); err != nil {
		fmt.Println("get website err:", err.Error())
		return website, pageData, err
	}
	if err = this.mongo.Collection("page").FindOne(
		context.TODO(),
		bson.M{"website_id": website.ID, "status": 1, "url": url}).Decode(&pageData); err != nil {
		fmt.Println("get page err:", err.Error())
		return website, pageData, err
	}
	return website, pageData, nil
}

func (this *websiteService) CreateWebsite(website *Website) (err error) {
	_, err = this.mongo.Collection("website").InsertOne(context.TODO(), website)
	if err != nil {
		return err
	}
	return nil
}

func (this *websiteService) UpdateWebsite(website *Website) (err error) {
	result, err := this.mongo.Collection("website").UpdateOne(
		context.TODO(),
		bson.M{
			"_id":     website.ID,
			"user_id": website.UserID,
		},
		bson.M{
			"$set": bson.M{
				"name":        website.Name,
				"category":    website.Category,
				"icon":        website.Icon,
				"keywords":    website.Keywords,
				"description": website.Description,
				"status":      website.Status,
			},
		})
	fmt.Println(result.MatchedCount, result.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

func (this *websiteService) DeleteWebsite(userID, websiteID primitive.ObjectID) (err error) {
	result, err := this.mongo.Collection("website").UpdateOne(context.TODO(), bson.M{"_id": websiteID, "user_id": userID}, bson.M{"status": 0})
	if err != nil {
		return err
	}
	fmt.Println(result.ModifiedCount)
	return nil
}
