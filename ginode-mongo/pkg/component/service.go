package component

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	ComponentService interface {
		GetComponentList(offset, limit int64, search Component) (component []Component, count int64, err error)
		GetComponent(id primitive.ObjectID) (component Component, err error)
		CreateComponent(component *Component) (err error)
		UpdateComponent(component Component) (err error)
		DeleteComponent(componentID primitive.ObjectID) (err error)
	}
)
type componentService struct {
	mongo *mongo.Database
}

func NewService(mongo *mongo.Database) ComponentService {
	return &componentService{
		mongo: mongo,
	}
}

func (u *componentService) GetComponentList(offset, limit int64, search Component) (components []Component, count int64, err error) {
	var cursor *mongo.Cursor
	if cursor, err = u.mongo.Collection("component").Find(
		context.TODO(),
		bson.M{"status": 1},
		options.Find().SetLimit(limit),
		options.Find().SetSkip(offset),
		options.Find().SetSort(bson.M{"sort": 1})); err != nil {
		return nil, 0, err
	}
	if err = cursor.All(context.TODO(), &components); err != nil {
		return nil, 0, err
	}
	//如需对数据迭代处理使用下面方法，不需要则直接使用上面方法
	//for cursor.Next(context.TODO()) {
	//	component := Component{}
	//	if err = cursor.Decode(&component); err != nil {
	//		fmt.Println(">>>>>>>>>.")
	//		return nil, 0, err
	//	}
	//	fmt.Println(component)
	//	components = append(components, component)
	//}
	fmt.Println(">>>>>>component list:", components)
	//查询集合里面有多少数据
	if count, err = u.mongo.Collection("component").CountDocuments(context.TODO(), bson.M{"status": 1}); err != nil {
		return nil, 0, err
	}

	fmt.Printf("Count:%d\n", count)
	return components, count, err
}

func (u *componentService) CreateComponent(component *Component) (err error) {
	_, err = u.mongo.Collection("component").InsertOne(context.TODO(), component)
	if err != nil {
		return err
	}
	return nil
}

func (u *componentService) GetComponent(id primitive.ObjectID) (component Component, err error) {
	if err = u.mongo.Collection("component").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&component); err != nil {
		return Component{}, err
	}
	return component, nil
}

func (u *componentService) UpdateComponent(component Component) (err error) {
	fmt.Printf(">>>>>>UpdateComponent:%+v", component)
	err = u.mongo.Collection("component").FindOneAndUpdate(
		context.TODO(),
		bson.D{{"_id", component.ID}},
		//bson.M{"$set": bson.M{
		//	"name": component.Name,
		//	"category": component.Category,
		//	"status": component.Status,
		//	"title_1": component.Title1,
		//	"title_2": component.Title2,
		//	"title_3": component.Title3,
		//	"description": component.Description,
		//	"background_img": component.BackgroundImg,
		//	"big_img": component.BigImg,
		//	"links": component.Links,
		//	"sort": component.Sort,
		//},
		// $set值直接传递结构体也是可以修改的，但是有个奇怪的现象是如果想要设置字段为空就不起作用，这是由于模型定义中的omitempty起来作用
		bson.M{"$set": component}).Decode(&component)
	if err != nil {
		return err
	}
	return nil
}

func (u *componentService) DeleteComponent(componentID primitive.ObjectID) (err error) {
	result, err := u.mongo.Collection("component").UpdateOne(
		context.TODO(),
		bson.D{{"_id", componentID}},
		bson.M{"$set": bson.M{
			"status": 0,
		}})
	if err != nil {
		return err
	}
	fmt.Println(result.ModifiedCount)
	return nil
}
