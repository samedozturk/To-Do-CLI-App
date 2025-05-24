package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Title      string             `bson:"title" json:"title"`
	Content    string             `bson:"content" json:"content"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	Done       bool               `bson:"done" json:"done"`
	ExpireDate time.Time          `bson:"expire_date" json:"expire_date"`
}
type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FullName  string             `bson:"name" json:"name"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type UserAccount struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId       primitive.ObjectID `bson:"user_id" json:"userId"`
	Email        string             `bson:"email" json:"email"`
	PhoneNumber  string             `bson:"phone_number" json:"phoneNumber"`
	PasswordHash string             `bson:"password_hash" json:"passwordHash"`
	Provider     string             `bson:"provider" json:"provider"`
	ProviderID   string             `bson:"provider_id" json:"providerID"`
}

type Token struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"userId"`
	Token     string             `bson:"token" json:"token"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
	ExpireAt  time.Time          `bson:"expire_at" json:"expireAt"`
	IsActive  bool               `bson:"is_active" json:"isActive"`
	UserAgent string             `bson:"user_agent,omitempty" json:"userAgent,omitempty"`
	IPAddress string             `bson:"ip_address,omitempty" json:"ipAddress,omitempty"`
}
