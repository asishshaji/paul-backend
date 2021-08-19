package repository

import (
	"backend/dto"
	"backend/model"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (uR *UserRepository) UpdateUserGenreScore(ctx context.Context, username, genre string, score int) error {
	res := uR.userCollection.FindOneAndUpdate(ctx, bson.M{"username": username},
		bson.M{"$set": bson.M{"genrescores." + genre: score}})

	if res.Err() == mongo.ErrNoDocuments {
		return errors.New("no user found")
	}

	return nil
}
func (uR *UserRepository) AddGenre(ctx context.Context, username, genre string) error {

	opts := options.FindOneAndUpdateOptions{}
	opts.SetUpsert(true)

	// addToSet : adds to array if not exists
	res := uR.userCollection.FindOneAndUpdate(ctx, bson.M{"username": username},
		bson.M{"$addToSet": bson.M{"genre": genre}}, &opts)

	if res.Err() == mongo.ErrNoDocuments {
		return errors.New("no user found")
	}
	return nil
}
