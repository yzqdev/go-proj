package component

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Component 网站组件
type Component struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`                       //bson是用来创建后返回，omitempty可选
	Name          string             `json:"name" bson:"name,omitempty"`                     //组件名
	Category      string             `json:"category" bson:"category,omitempty"`             //分类
	Title1        string             `json:"title_1,omitempty" bson:"title_1"`               //一级标题
	Title2        string             `json:"title_2,omitempty" bson:"title_2"`               //二级标题
	Title3        string             `json:"title_3,omitempty" bson:"title_3"`               //三级标题
	Description   string             `json:"description,omitempty" bson:"description"`       //文字描述
	BackgroundImg string             `json:"background_img,omitempty" bson:"background_img"` //背景图
	BigImg        string             `json:"big_img,omitempty" bson:"big_img"`               //展示大图
	Elements      []Element          `json:"elements" bson:"elements,omitempty"`             //多个同类型元素列表展示
	Links         []Link             `json:"links,omitempty" bson:"links,omitempty"`         //链接或按钮
	Extras        []Component        `json:"extras,omitempty" bson:"extras,omitempty"`       //其他补充内容
	Status        int                `json:"status,omitempty" bson:"status,omitempty"`       //可用状态
	Sort          int                `json:"sort,omitempty" bson:"sort,omitempty"`           //排序
	CreatedAt     *time.Time         `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

//Element 组件中包含的同类型元素
type Element struct {
	Icon        string `json:"icon,omitempty" bson:"icon,omitempty"` //小图标
	Img         string `json:"img,omitempty" bson:"img,omitempty"`   //大图
	Title1      string `json:"title_1,omitempty" bson:"title_1,omitempty"`
	Title2      string `json:"title_2,omitempty" bson:"title_2,omitempty"`
	Title3      string `json:"title_3,omitempty" bson:"title_3,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Link        string `json:"link,omitempty" bson:"link,omitempty"` //跳转链接，也可以用作视频链接
}

//Link 组件包含的链接
type Link struct {
	Title string `json:"title" bson:"title,omitempty"`
	URL   string `json:"url" bson:"url,omitempty"`
}
