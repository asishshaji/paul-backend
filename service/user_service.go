package service

import (
	"backend/dto"
	"backend/model"
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

func (uS UserService) CreateUser(ctx context.Context, user dto.UserRegDto) (model.User, error) {

	// Check if the user with username exists
	err := uS.userRepo.CheckIfUserWithUsernameExists(ctx, user.Username)
	if err != mongo.ErrNoDocuments {
		log.Println(err)
		return model.User{}, errors.New("User already exists")
	}

	us := &model.User{
		Username:   user.Username,
		Password:   user.Password,
		GenreScore: map[string]int{},
		Genre:      []string{},
	}

	err = uS.userRepo.CreateUser(ctx, us)
	if err != nil {
		log.Println(err)
		return model.User{}, errors.New("Error creating user")
	}

	return *us, nil
}

func (uS UserService) LoginUser(ctx context.Context, user dto.UserRegDto) error {

	err := uS.userRepo.CheckIfUserWithNameAndPasswordExists(ctx, user)

	return err
}

func (uS UserService) AddGenre(ctx context.Context, genre, username string) error {
	log.Println(username)
	err := uS.userRepo.AddGenre(ctx, username, genre)
	if err != nil {
		return err
	}
	return nil

}
