package service

import (
	"backend/model"
	"context"
)

type QuizService struct {
}

func NewQuizService() IQuizService {
	return &QuizService{}
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
