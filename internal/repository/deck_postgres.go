package repository

import (
	"database/sql"
	"errors"

	"github.com/J-V-S-C/flashCards-go/internal/models"
)

type deckPostgres struct {
	db *sql.DB
}

func NewDeckPostgres(db *sql.DB) *deckPostgres {
	return &deckPostgres{db: db}
}

func (receiver *deckPostgres) AddDeck(deck models.Deck) (int, error) {
	var deckId int
	if deck.Name == "" {
		return 0, errors.New("please provide a valid name")
	}

	query := `Insert INTO deck (name, total_cards, next_review_at) VALUES ($1, $2, $3) RETURNING id`

	err := receiver.db.QueryRow(query, deck.Name, deck.TotalCards, deck.NextReviewAt).Scan(&deckId)

	if err != nil {
		return 0, err
	}
	return deckId, nil
}
