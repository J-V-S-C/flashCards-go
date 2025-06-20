package contracts

import "github.com/J-V-S-C/flashCards-go/models"

type Deck interface {
	GetAll() ([]models.Deck, error)
	GetById(deckId int) (models.Deck, error)
	AddDeck(deck models.Deck) (int, error)
	UpdateDeck(deck models.Deck, deckId int) error
	DeleteDeck(deckId int) error
}
