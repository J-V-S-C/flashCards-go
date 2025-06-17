package repository

import (
	"database/sql"
	"errors"

	"github.com/J-V-S-C/flashCards-go/internal/models"
)

type FlashcardPostgres struct {
	db *sql.DB
}

func NewFlashcardPostgres(db *sql.DB) *FlashcardPostgres {
	return &FlashcardPostgres{db: db}
}

// receiver is equivalent to "this"
func (receiver *FlashcardPostgres) AddFlashcard(flashcard models.Flashcard) (int, error) {
	var flashcardId int

	if flashcard.Name == "" && flashcard.DeckId <= 0 {
		return 0, errors.New("please provide valid flashcard name from a valid deck")
	}

	query := `INSERT INTO flashcard (name, deck_id, message, next_review_at ,ease_factor) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	row := receiver.db.QueryRow(query, flashcard.Name, flashcard.DeckId, flashcard.Message, flashcard.NextReviewAt, flashcard.EaseFactor)

	err := row.Scan(&flashcardId)
	if err != nil {
		return 0, err
	}
	return flashcardId, nil
}

func (receiver *FlashcardPostgres) GetAllFlashcards() ([]models.Flashcard, error) {
	query := `SELECT id, name, deck_id, message, next_review_at, ease_factor FROM flashcard`

	res, err := receiver.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	selectedFlashcard := []models.Flashcard{}

	for res.Next() {
		flashcard := models.Flashcard{}

		err := res.Scan(&flashcard.Id, &flashcard.Name, &flashcard.DeckId, &flashcard.Message, &flashcard.NextReviewAt, &flashcard.EaseFactor)
		if err != nil {
			return nil, err
		}

		selectedFlashcard = append(selectedFlashcard, flashcard)
	}
	return selectedFlashcard, nil
}

func (receiver *FlashcardPostgres) GetById(flashcardId int) (models.Flashcard, error) {
	flashcard := models.Flashcard{}

	query := `SELECT id, name, deck_id, message, next_review_at FROM flashcard WHERE id=$1`
	err := receiver.db.QueryRow(query, flashcardId).Scan(&flashcard.Id, &flashcard.Name, &flashcard.DeckId, &flashcard.Message, &flashcard.NextReviewAt, &flashcard.EaseFactor)
	if err != nil {
		return models.Flashcard{}, err
	}

	return flashcard, nil
}

func (receiver *FlashcardPostgres) UpdateFlashcard(newFlashcard models.Flashcard, flashcardId int) error {
	var scannerId int
	query := `SELECT id FROM flashcard WHERE id=$1`
	err := receiver.db.QueryRow(query, flashcardId).Scan(&scannerId)
	if err != nil {
		return err
	}

	if newFlashcard.Name == "" && newFlashcard.DeckId <= 0 {
		return errors.New("please provide valid flashcard name from a valid deck")
	}

	query = `UPDATE flashcard SET name=$1, deck_id=$2, message=$3, next_review_at=$4, ease_factor=$5 WHERE id=$6`
	_, err = receiver.db.Exec(query, newFlashcard.Name, newFlashcard.DeckId, newFlashcard.Message, newFlashcard.NextReviewAt, newFlashcard.EaseFactor, flashcardId)
	if err != nil {
		return err
	}

	return nil

}

func (receiver *FlashcardPostgres) DeleteFlashcard(flashcardId int) error {
	var scannerId int
	query := `SELECT id FROM flashcard WHERE id=$1`
	err := receiver.db.QueryRow(query, flashcardId).Scan(&scannerId)
	if err != nil {
		return err
	}

	query = `DELETE FROM flashcard WHERE id=$1`
	_, err = receiver.db.Exec(query, flashcardId)
	if err != nil {
		return err
	}
	return nil
}
