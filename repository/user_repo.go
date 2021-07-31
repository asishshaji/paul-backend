package repository

import (
	"backend/dto"
	"backend/model"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, userCollectionPath string) IUserRepository {
	return &UserRepository{
		userCollection: db.Collection(userCollectionPath),
	}
}

func (uR *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	res, err := uR.userCollection.InsertOne(ctx, user)

	if err != nil {
		return err
	}
	log.Println("Created new User with id :", res.InsertedID)
	return nil
}

func (uR *UserRepository) CheckIfUserWithUsernameExists(ctx context.Context, username string) error {

	return uR.userCollection.FindOne(ctx, bson.M{
		"username": username,
	}).Err()

}

func (uR *UserRepository) CheckIfUserWithNameAndPasswordExists(ctx context.Context, user dto.UserRegDto) error {
	filter := bson.M{
		"username": user.Username,
		"password": user.Password,
	}
	res := uR.userCollection.FindOne(ctx, filter)

	if res.Err() == mongo.ErrNoDocuments {
		return errors.New("invalid credentials")
	}

	return nil

}
