package page

import (
	"time"

	"github.com/yuedun/ginode-mongo/pkg/component"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Page 页面数据结构
type Page struct {
	ID          primitive.ObjectID    `json:"_id" bson:"_id,omitempty"` //bson是用来创建后返回，omitempty是可选
	WebsiteID   primitive.ObjectID    `json:"website_id" bson:"website_id,omitempty"`
	Title       string                `json:"title" bson:"title,omitempty"`             //网站名 TDK title
	Keywords    string                `json:"keywords" bson:"keywords,omitempty"`       //关键字 TDK keywords
	Description string                `json:"description" bson:"description,omitempty"` //描述 TDK description
	URL         string                `json:"url" bson:"url,omitempty"`                 //网站地址
	Components  []component.Component `json:"components" bson:"components,omitempty"`   //包含组件
	Status      int                   `json:"status" bson:"status,omitempty"`           //状态：0不可用，1可用
	CreatedAt   time.Time             `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time             `json:"updated_at" bson:"updated_at,omitempty"`
}
