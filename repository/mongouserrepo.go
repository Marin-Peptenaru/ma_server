package repository

import (
	"books/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	users mongo.Collection
}

func NewUserRepository(conn *mongo.Client) UserRepository {
	return &mongoUserRepository{
		users: *conn.Database("ma").Collection("users"),
	}

}

func (u *mongoUserRepository) GetUser(username string) (domain.User, error) {
	usernameFilter := bson.D{{"username", username}}
	queryResult := u.users.FindOne(context.TODO(), usernameFilter)
	user := domain.User{}

	if queryResult.Err() != nil {
		return user, errors.New("error fetchig user: " + queryResult.Err().Error())
	}


	decodeErr := queryResult.Decode(&user)

	if decodeErr != nil {
		return user, errors.New("error decoding fetched user: " + decodeErr.Error())
	}

	return user, nil

}