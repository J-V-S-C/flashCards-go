package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/J-V-S-C/flashCards-go/utils"
)

// @Summary      Adiciona um novo deck
// @Description  Cria um novo deck com os dados fornecidos
// @Tags         decks
// @Accept       json
// @Produce      json
// @Param        deck  body      models.Deck  true  "Deck a ser criado"
// @Success      201   {object}  models.Deck
// @Failure      500   {object}  utils.ErrResponseJSON
// @Router       /decks/ [post]
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
	deck.Id = id
	utils.WriteJSON(w, deck, http.StatusCreated)
}

// @Summary      Lista todos os decks
// @Description  Retorna uma lista com todos os decks
// @Tags         decks
// @Produce      json
// @Success      200  {array}   models.Deck
// @Failure      500  {object}  utils.ErrResponseJSON
// @Router       /decks/ [get]
func (h *DeckHandler) getAll(w http.ResponseWriter, r *http.Request) {
	decks, err := h.services.GetAll()
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, decks, http.StatusOK)
}

// @Summary      Busca deck por ID
// @Description  Retorna um deck pelo seu ID
// @Tags         decks
// @Produce      json
// @Param        id   path      int  true  "ID do Deck"
// @Success      200  {object}  models.Deck
// @Failure      400  {object}  utils.ErrResponseJSON
// @Failure      500  {object}  utils.ErrResponseJSON
// @Router       /decks/{id} [get]
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

// @Summary      Atualiza um deck
// @Description  Atualiza os dados de um deck existente
// @Tags         decks
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "ID do Deck"
// @Param        deck  body      models.Deck true  "Novos dados do deck"
// @Success      200   {object}  models.Deck
// @Failure      400   {object}  utils.ErrResponseJSON
// @Failure      500   {object}  utils.ErrResponseJSON
// @Router       /decks/{id} [put]
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

// @Summary      Deleta um deck
// @Description  Remove um deck pelo ID
// @Tags         decks
// @Produce      json
// @Param        id   path      int  true  "ID do Deck"
// @Success      200  {string}  string  "Mensagem de sucesso"
// @Failure      400  {object}  utils.ErrResponseJSON
// @Failure      500  {object}  utils.ErrResponseJSON
// @Router       /decks/{id} [delete]
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
