package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/J-V-S-C/flashCards-go/internal/models"
)

func DecodeFlashcardJSON(req *http.Request) (models.Flashcard, error) {
	flashcard := models.Flashcard{}

	err := json.NewDecoder(req.Body).Decode(&flashcard)
	if err != nil {
		return models.Flashcard{}, err
	}
	return flashcard, nil
}

func DecodeDeckJSON(req *http.Request) (models.Deck, error) {
	deck := models.Deck{}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&deck)
	if err != nil {
		return deck, err
	}
	return deck, nil
}
