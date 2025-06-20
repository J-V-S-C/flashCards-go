package handlers

import (
	"net/http"

	"github.com/J-V-S-C/flashCards-go/internal/contracts"
	"github.com/J-V-S-C/flashCards-go/internal/service"
)

type FlashcardHandler struct {
	services contracts.Flashcard
}

type DeckHandler struct {
	services contracts.Deck
}

type Handler struct {
	Flashcard *FlashcardHandler
	Deck      *DeckHandler
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Flashcard: &FlashcardHandler{services: services.Flashcard},
		Deck: &DeckHandler{services: services.Deck},
	}
}

func (h *FlashcardHandler) InitFlashcardRoutes() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("POST /", h.addFlashcard)
	r.HandleFunc("GET /", h.getAll)
	r.HandleFunc("GET /{id}", h.getById)
	r.HandleFunc("PUT /{id}", h.UpdateFlashcard)
	r.HandleFunc("DELETE /{id}", h.DeleteFlashcard)
	return r
}

func (h *DeckHandler) InitDeckRoutes() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("POST /", h.addDeck)
	r.HandleFunc("GET /", h.getAll)
	r.HandleFunc("GET /{id}", h.getById)
	r.HandleFunc("PUT /{id}", h.UpdateDeck)
	r.HandleFunc("DELETE /{id}", h.DeleteDeck)
	return r
}

func (h *Handler) InitRoutes() http.Handler {
	r := http.NewServeMux()
	r.Handle("/flashcards/", h.Flashcard.InitFlashcardRoutes())
	r.Handle("/decks/", h.Deck.InitDeckRoutes())
	return r
}
