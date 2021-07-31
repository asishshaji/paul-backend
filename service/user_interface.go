package service

import (
	"backend/dto"
	"context"
)

type IUserService interface {
	CreateUser(ctx context.Context, user dto.UserRegDto) error
	LoginUser(ctx context.Context, user dto.UserRegDto) error
	AddGenre(ctx context.Context, genre, username string) error
}
