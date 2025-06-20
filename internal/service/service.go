package service

import (
	"github.com/J-V-S-C/flashCards-go/internal/contracts"
	"github.com/J-V-S-C/flashCards-go/internal/repository"
)

type Service struct {
	Flashcard contracts.Flashcard
	Deck      contracts.Deck
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Flashcard: NewFlashcardListService(repo.Flashcard),
		Deck:      NewDeckListService(repo.Deck),
	}
}
