package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/J-V-S-C/flashCards-go/utils"
)

// @Summary      Adiciona um novo flashcard
// @Description  Cria um novo flashcard com os dados fornecidos
// @Tags         flashcards
// @Accept       json
// @Produce      json
// @Param        flashcard  body      models.Flashcard  true  "Flashcard a ser criado"
// @Success      201        {object}  models.Flashcard
// @Failure      500        {object}  utils.ErrResponseJSON
// @Router       /flashcards/ [post]
func (h *FlashcardHandler) addFlashcard(w http.ResponseWriter, r *http.Request) {
	flashcard, err := utils.DecodeFlashcardJSON(r)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	id, err := h.services.AddFlashcard(flashcard)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	flashcard.Id = id
	utils.WriteJSON(w, flashcard, http.StatusCreated)
}

// @Summary      Lista todos os flashcards
// @Description  Retorna todos os flashcards disponíveis
// @Tags         flashcards
// @Produce      json
// @Success      200  {array}   models.Flashcard
// @Failure      500  {object}  utils.ErrResponseJSON
// @Router       /flashcards/ [get]
func (h *FlashcardHandler) getAll(w http.ResponseWriter, r *http.Request) {
	flashcards, err := h.services.GetAll()
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, flashcards, http.StatusOK)
}

// @Summary      Busca flashcard por ID
// @Description  Retorna um flashcard com base no ID fornecido
// @Tags         flashcards
// @Produce      json
// @Param        id   path      int  true  "ID do flashcard"
// @Success      200  {object}  models.Flashcard
// @Failure      400  {object}  utils.ErrResponseJSON
// @Failure      500  {object}  utils.ErrResponseJSON
// @Router       /flashcards/{id} [get]
func (h *FlashcardHandler) getById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	flashcard, err := h.services.GetById(id)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, flashcard, http.StatusOK)
}

// @Summary      Atualiza flashcard
// @Description  Atualiza um flashcard existente pelo ID
// @Tags         flashcards
// @Accept       json
// @Produce      json
// @Param        id         path      int               true  "ID do flashcard"
// @Param        flashcard  body      models.Flashcard  true  "Dados atualizados"
// @Success      200        {object}  models.Flashcard
// @Failure      400        {object}  utils.ErrResponseJSON
// @Failure      500        {object}  utils.ErrResponseJSON
// @Router       /flashcards/{id} [put]
func (h *FlashcardHandler) UpdateFlashcard(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	newFlashcard, err := utils.DecodeFlashcardJSON(r)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	err = h.services.UpdateFlashcard(newFlashcard, id)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, newFlashcard, http.StatusOK)
}

// @Summary      Remove flashcard
// @Description  Remove um flashcard com base no ID fornecido
// @Tags         flashcards
// @Produce      json
// @Param        id   path      int  true  "ID do flashcard"
// @Success      200  {object}  utils.MsgResponseJSON
// @Failure      400  {object}  utils.ErrResponseJSON
// @Failure      500  {object}  utils.ErrResponseJSON
// @Router       /flashcards/{id} [delete]
func (h *FlashcardHandler) DeleteFlashcard(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	err = h.services.DeleteFlashcard(id)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	responseString := fmt.Sprintf("%d deleted", id)
	utils.RespondWithMessage(w, responseString, http.StatusOK)
}
