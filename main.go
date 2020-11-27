package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alextanhongpin/go-mongo/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDB() *mongo.Client {
	var (
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		db   = os.Getenv("DB_NAME")
		uri  = fmt.Sprintf("mongodb://%s:%s/%s", host, port, db)
	)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(options.Credential{
		//AuthSource: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}))

	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	client := setupDB()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	userStore := domain.NewUserStore(client)
	var createUser bool
	if createUser {
		u, err := userStore.Create(context.Background(), domain.CreateUserParams{
			Name:  "John Doe",
			Email: "john.doe@mail.com",
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("user created: %#v", u)
	}
	users, err := userStore.FindAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("found users: %#v", len(users))
	for _, u := range users {
		log.Printf("%#v\n", u)
	}

	uu, err := userStore.FindOne(context.Background(), users[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("found user: %#v", uu)
}
