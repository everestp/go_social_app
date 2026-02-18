package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string             `json:"content" bson:"content"`
	Sender  string             `json:"sender" bson:"sender"`
	Receiver string             `json:"recever" bson:"receiver"`
}

// interfaces
type SendMessage struct {
	Content string `json:"content" bson:"content"  validate:"required,min=5"`
	Sender  string `json:"sender" bson:"sender"  validate:"required"`
	Recever string `json:"recever" bson:"recever"  validate:"required"`
}