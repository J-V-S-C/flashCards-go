package models

import "time"

type Flashcard struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	DeckId       int       `json:"deck_id"`
	Message      string    `json:"message"`
	NextReviewAt time.Time `json:"next_review_at"`
	EaseFactor   float64   `json:"ease_factor"`
}
