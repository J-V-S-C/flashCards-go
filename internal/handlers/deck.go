package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/J-V-S-C/flashCards-go/utils"
)

func (h *DeckHandler) addDeck(w http.ResponseWriter, r *http.Request) {
	deck, err := utils.DecodeDeckJSON(r)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	id, err := h.services.AddDeck(deck)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	// By default "id" is a zero value, but we need to show the user what the sql "id" of updated Deck is
	deck.Id = id
	utils.WriteJSON(w, deck, http.StatusCreated)
}

func (h *DeckHandler) getAll(w http.ResponseWriter, r *http.Request) {
	decks, err := h.services.GetAll()
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, decks, http.StatusOK)
}

func (h *DeckHandler) getById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	deck, err := h.services.GetById(id)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, deck, http.StatusOK)
}

func (h *DeckHandler) UpdateDeck(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	newDeck, err := utils.DecodeDeckJSON(r)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	err = h.services.UpdateDeck(newDeck, id)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	newDeck.Id = id
	utils.WriteJSON(w, newDeck, http.StatusOK)
}

func (h *DeckHandler) DeleteDeck(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	err = h.services.DeleteDeck(id)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	responseString := fmt.Sprintf("%d deleted", id)
	utils.RespondWithMessage(w, responseString, 200)
}
