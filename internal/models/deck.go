package models

import "time"

type Deck struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	TotalCards   int       `json:"total_cards"`
	NextReviewAt time.Time `json:"next_review_at"`
}
