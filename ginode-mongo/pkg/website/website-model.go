package website

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Website struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`                 //bson是用来创建后返回，omitempty是可选
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`         //用户id
	Name        string             `json:"name" bson:"name,omitempty"`               //网站名 TDK title
	Description string             `json:"description" bson:"description,omitempty"` //描述 TDK description
	Keywords    string             `json:"keywords" bson:"keywords,omitempty"`       //关键字 TDK keywords
	Icon        string             `json:"icon" bson:"icon,omitempty"`               //图标
	Category    string             `json:"category" bson:"category,omitempty"`       //网站分类
	Status      int                `json:"status" bson:"status,omitempty"`           //状态
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
