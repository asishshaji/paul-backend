package repository

import (
	"backend/entity"
	"context"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	CheckIfUserWithUsernameExists(ctx context.Context, username string) error
	// FindAll() ([]entity.User, error)
}
