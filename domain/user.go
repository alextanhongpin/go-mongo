package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
