package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"` //bson是用来创建后返回，omitempty是创建时自动创建
	Mobile      string             `json:"mobile" bson:"mobile,omitempty"`
	Username    string             `json:"username" bson:"username,omitempty"`
	Password    string             `json:"password" bson:"password,omitempty"`
	Gender      string             `json:"gender" bson:"gender,omitempty"`
	Avatar      string             `json:"avatar" bson:"avatar,omitempty"`
	Addr        string             `json:"addr" bson:"addr,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Status      int                `json:"status" bson:"status,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
