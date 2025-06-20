package service

import (
	"github.com/J-V-S-C/flashCards-go/internal/contracts"
	"github.com/J-V-S-C/flashCards-go/models"
)

type FlashcardService struct {
	repo contracts.Flashcard
}

func NewFlashcardListService(repo contracts.Flashcard) *FlashcardService {
	return &FlashcardService{repo: repo}
}

func (s *FlashcardService) AddFlashcard(flashcardToAdd models.Flashcard) (int, error) {
	return s.repo.AddFlashcard(flashcardToAdd)
}

func (s *FlashcardService) GetAll() ([]models.Flashcard, error) {
	return s.repo.GetAll()
}

func (s *FlashcardService) GetById(flashcardId int) (models.Flashcard, error) {
	return s.repo.GetById(flashcardId)
}

func (s *FlashcardService) UpdateFlashcard(newFlashcard models.Flashcard, flashcardId int) error {
	return s.repo.UpdateFlashcard(newFlashcard, flashcardId)
}

func (s *FlashcardService) DeleteFlashcard(flashcardId int) error {
	return s.repo.DeleteFlashcard(flashcardId)
}
