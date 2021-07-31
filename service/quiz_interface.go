package service

import (
	"backend/model"
	"context"
)

type IQuizService interface {
	GetQuizBasedOnGenre(ctx context.Context, genre string) []model.Quiz
	SaveScore(ctx context.Context, score int, username, genre string) error
}
