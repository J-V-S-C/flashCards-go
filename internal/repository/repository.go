package repository

import (
	"database/sql"

	"github.com/J-V-S-C/flashCards-go/internal/contracts"
)

type Repository struct {
	Flashcard contracts.Flashcard
	Deck      contracts.Deck
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Flashcard: NewFlashcardPostgres(db),
		Deck:      NewDeckPostgres(db),
	}
}
