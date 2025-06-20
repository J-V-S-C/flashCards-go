package repository

import (
	"database/sql"
	"errors"

	"github.com/J-V-S-C/flashCards-go/models"
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

func (receiver *deckPostgres) GetAll() ([]models.Deck, error) {
	query := `SELECT id, name, total_cards, next_review_at FROM deck`
	res, err := receiver.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	selectedDeck := []models.Deck{}

	for res.Next() {
		deck := models.Deck{}
		err := res.Scan(&deck.Id, &deck.Name, &deck.TotalCards, &deck.NextReviewAt)
		if err != nil {
			return nil, err
		}

		selectedDeck = append(selectedDeck, deck)
	}
	return selectedDeck, nil
}

func (receiver *deckPostgres) GetById(deckId int) (models.Deck, error) {
	deck := models.Deck{}
	query := `SELECT id,name, total_cards, next_review_at FROM deck WHERE id=$1`
	err := receiver.db.QueryRow(query, deckId).Scan(&deck.Id, &deck.Name, &deck.TotalCards, &deck.NextReviewAt)
	if err != nil {
		return models.Deck{}, err
	}

	return deck, nil

}

func (receiver *deckPostgres) UpdateDeck(newDeck models.Deck, deckId int) error {

	if newDeck.Name == "" {
		return errors.New("Please provide valid deck name")
	}
	query := `UPDATE deck SET name=$1, total_cards=$2, next_review_at=$3 WHERE id=$4`
	res, err := receiver.db.Exec(query, newDeck.Name, newDeck.TotalCards, newDeck.NextReviewAt, deckId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Deck not found")
	}
	return nil

}

func (receiver *deckPostgres) DeleteDeck(deckId int) error {
	query := `DELETE FROM deck WHERE id=$1`
	res, err := receiver.db.Exec(query, deckId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("Deck not found")
	}
	return nil
}
