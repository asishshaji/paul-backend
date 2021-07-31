package service

import (
	"backend/model"
	"context"
)

type IQuizService interface {
	GetQuizBasedOnGenre(ctx context.Context, genre string) []model.Quiz
}
