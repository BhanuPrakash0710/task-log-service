package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type RegisterInputUser struct {
	Email    string `json:"email" bson:"email"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}

type LoginInputUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type User struct {
	ID       string `json:"id"  bson:"_id"`
	Email    string `json:"email" bson:"email"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	//UserId   string `json:"userId" bson:"userId"`
}

func CreateUser(user *User) error {
	collection := mongoClient.Database(DatabaseName).Collection(UsercollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return err
	}

	fmt.Println("Created a record with id:", result.InsertedID)
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	collection := mongoClient.Database(DatabaseName).Collection(UsercollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		fmt.Println("Error finding user:", err)
		return nil, err
	}

	return &user, nil
}
