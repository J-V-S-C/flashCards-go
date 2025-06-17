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

func (receiver *deckPostgres) GetAllDecks() ([]models.Deck, error) {
	query := `SELECT id, name, total_cards, next_review_at FROM deck`
	res, err := receiver.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	selectedDeck := []models.Deck{}

	for res.Next() {
		deck := models.Deck{}
		err := res.Scan(&deck.Id, &deck.Name, &deck.NextReviewAt, &deck.TotalCards)
		if err != nil {
			return nil, err
		}

		selectedDeck = append(selectedDeck, deck)
	}
	return selectedDeck, nil
}
