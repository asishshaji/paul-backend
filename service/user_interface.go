package service

import (
	"backend/dto"
	"backend/model"
	"context"
)

type IUserService interface {
	CreateUser(ctx context.Context, user dto.UserRegDto) (model.User, error)
	LoginUser(ctx context.Context, user dto.UserRegDto) error
	AddGenre(ctx context.Context, genre, username string) error
}
