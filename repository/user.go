package repository

import (
	"backend/entity"
	"context"
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

func (uR *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
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
