package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/J-V-S-C/flashCards-go/models"
)

type ErrResponseJSON struct {
	ErrMsg string `json:"error"`
}

type MsgResponseJSON struct {
	Msg string `json:"message"`
}

func RespondWithMessage(w http.ResponseWriter, receivedMsg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	msg := MsgResponseJSON{Msg: receivedMsg}
	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		log.Println("error enconding json:", err)
	}
}

func RespondWithError(w http.ResponseWriter, receivedError error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errMsg := ErrResponseJSON{ErrMsg: receivedError.Error()}
	err := json.NewEncoder(w).Encode(errMsg)
	if err != nil {
		log.Println("error enconding json:", err)
	}
}

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

func WriteJSON(w http.ResponseWriter, data any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
