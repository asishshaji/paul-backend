package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
	"context"
	"log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(uR repository.IUserRepository) IUserService {
	return UserService{
		userRepo: uR,
	}
}

func (uS UserService) CreateUser(ctx context.Context, user dto.UserRegDto) error {

	// Check if the user with username exists
	err := uS.userRepo.CheckIfUserWithUsernameExists(ctx, user.Username)
	if err != mongo.ErrNoDocuments {
		log.Println(err)
		return errors.New("User already exists")
	}

	// Create user here
	// convert user dto to user entity
	err = uS.userRepo.CreateUser(ctx, &entity.User{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		log.Println(err)
		return errors.New("Error creating user")
	}

	return nil
}

func (uS UserService) LoginUser(ctx context.Context, username, password string) error {
	return nil
}
