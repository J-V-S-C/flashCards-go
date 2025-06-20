package contracts

import "github.com/J-V-S-C/flashCards-go/models"

type Flashcard interface {
	GetAll() ([]models.Flashcard, error)
	GetById(flashcardId int) (models.Flashcard, error)
	AddFlashcard(flashcard models.Flashcard) (int, error)
	UpdateFlashcard(flashcard models.Flashcard, flashcardId int) error
	DeleteFlashcard(flashcardId int) error
}
