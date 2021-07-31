package service

import (
	"backend/model"
	"backend/repository"
	"context"
)

type QuizService struct {
	userRepo repository.IUserRepository
}

func NewQuizService(
	userRepo repository.IUserRepository) IQuizService {
	return &QuizService{userRepo: userRepo}
}

var QuizQuestions = []model.Quiz{
	{
		Type:     "other",
		Options:  []string{"A", "B", "C", "D"},
		Question: "Excepteur eiusmod minim pariatur occaecat nisi exercitation nulla aute quis nostrud eiusmod consectetur tempor?",
		Answer:   "A",
	},
	{
		Type:     "other",
		Options:  []string{"A", "B", "C", "D"},
		Question: "Excepteur eiusmod minim pariatur occaecat nisi exercitation nulla aute quis nostrud eiusmod consectetur tempor?",
		Answer:   "B",
	},
	{
		Type:     "other",
		Options:  []string{"A", "B", "C", "D"},
		Question: "Excepteur eiusmod minim pariatur occaecat nisi exercitation nulla aute quis nostrud eiusmod consectetur tempor?",
		Answer:   "D",
	},
}

func (qS QuizService) GetQuizBasedOnGenre(ctx context.Context, genre string) []model.Quiz {

	return QuizQuestions
}

func (qS QuizService) SaveScore(ctx context.Context, score int, username, genre string) error {
	err := qS.userRepo.UpdateUserGenreScore(ctx, username, genre, score)
	return err
}
