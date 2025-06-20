package service

import (
	"github.com/J-V-S-C/flashCards-go/internal/repository"
	"github.com/J-V-S-C/flashCards-go/models"
)

type DeckService struct {
	repo repository.Deck
}

func NewDeckListService(repo repository.Deck) *DeckService {
	return &DeckService{repo: repo}
}

func (s *DeckService) AddDeck(DeckToAdd models.Deck) (int, error) {
	return s.repo.AddDeck(DeckToAdd)
}

func (s *DeckService) GetAll() ([]models.Deck, error) {
	return s.repo.GetAll()
}

func (s *DeckService) GetById(DeckId int) (models.Deck, error) {
	return s.repo.GetById(DeckId)
}

func (s *DeckService) UpdateDeck(newDeck models.Deck, DeckId int) error {
	return s.repo.UpdateDeck(newDeck, DeckId)
}

func (s *DeckService) DeleteDeck(DeckId int) error {
	return s.repo.DeleteDeck(DeckId)
}
