package repository

import (
	"database/sql"
	"github.com/J-V-S-C/flashCards-go/internal/models"
)

type Flashcard interface {
	GetAll() ([]models.Flashcard, error)
	GetById(flashcardId int) (models.Flashcard, error)
	AddFlashcard(flashcard models.Flashcard) (int, error)
	UpdateFlashcard(flashcard models.Flashcard, flashcardId int) error
	DeleteFlashcard(flashcardId int) error
}
type Deck interface {
	GetAll() ([]models.Deck, error)
	GetById(deckId int) (models.Deck, error)
	AddDeck(deck models.Deck) (int, error)
	UpdateDeck(deck models.Deck, deckId int) error
	DeleteDeck(deckId int) error
}

type Repository struct {
	Flashcard
	Deck
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Flashcard: NewFlashcardPostgres(db),
		Deck:      NewDeckProstgres(db),
	}
}
