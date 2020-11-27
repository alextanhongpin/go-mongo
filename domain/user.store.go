package domain

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var DATABASE = os.Getenv("DB_NAME")

const UserCollection = "users"

type UserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserStore(c *mongo.Client) *UserStore {
	return &UserStore{
		client:     c,
		collection: c.Database(DATABASE).Collection(UserCollection),
	}
}

func (u *UserStore) FindOne(ctx context.Context, id primitive.ObjectID) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user User
	filter := bson.M{
		"_id": id,
	}
	if err := u.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserStore) FindAll(ctx context.Context) ([]User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var users []User
	cur, err := u.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

type CreateUserParams struct {
	Name  string
	Email string
}

func (u *UserStore) Create(ctx context.Context, params CreateUserParams) (User, error) {
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.collection.InsertOne(ctx, bson.M{
		"name":  params.Name,
		"email": params.Email,
	})
	if err != nil {
		return user, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	user.Name = params.Name
	user.Email = params.Email

	return user, nil
}
