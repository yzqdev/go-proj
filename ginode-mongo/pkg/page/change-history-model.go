package page

import (
	"time"

	"github.com/yuedun/ginode-mongo/pkg/component"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChangeHistory 修改历史，便于回滚
type ChangeHistory struct {
	ID          primitive.ObjectID    `json:"_id" bson:"_id,omitempty"` //bson是用来创建后返回，omitempty是可选
	PageID      primitive.ObjectID    `json:"page_id" bson:"page_id"`
	Title       string                `json:"title" bson:"title,omitempty"`             //网站名 TDK title
	Keywords    string                `json:"keywords" bson:"keywords,omitempty"`       //关键字 TDK keywords
	Description string                `json:"description" bson:"description,omitempty"` //描述 TDK description
	URL         string                `json:"url" bson:"url,omitempty"`                 //网站地址
	Components  []component.Component `json:"components" bson:"components,omitempty"`   //包含组件
	Status      int                   `json:"status" bson:"status,omitempty"`           //状态：0不可用，1可用
	CreatedAt   time.Time             `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time             `json:"updated_at" bson:"updated_at,omitempty"`
}
